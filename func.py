def get_periods(period, ts):
    "返回每个short_len bar的最后1分钟的时间区间"
    end_hour_list = list()
    for i in range(24):
        if (i+1) % period == 0:
            end_hour_list.append(i)
    tp = ts2tp(ts)
    tp_str = str(tp.date())
    endperiod_list = list()
    for i in end_hour_list:
        ts_start =  timestr2timestamp(tp_str + " %s:59:00" % i)
        ts_end = ts_start + 60
        endperiod_list.append([ts_start, ts_end])
    return endperiod_list

def isperiodend(ts, endperiod_list):
    # 返回是否是最后的1min，以及不能加仓区间的结束时间
    for terms in endperiod_list:
        if ts > terms[0] and ts < terms[1]:
            return True, terms[1]
    return False, None




class OrderInfo:
    def __init__(self, symbol, max_queue_num=1440, logger=None, api=None, is_realtime=0):
        self.deque = deque()
        self.init()
        self.symbol = symbol
        self.max_queue_num = max_queue_num
        # self.target_pos = target_pos  # 这个会变，运行一段时间看看效果
        self.logger = logger
        self.api = api
        self.is_realtime = is_realtime
    
    def __len__(self):
        return len(self.deque)

    def init(self):
        self.deque.clear()
        self.max_atr = 0
        self.hh = float("-inf")  #
        self.ll = float("inf")
        self.out_price = 0.  # 出场价格
    
    @timmer
    def append(self, order):
        self.deque.append(order)

        # 开多：买入开多（side 填写 buy； posSide 填写 long ）
        # 开空：卖出开空（side 填写 sell； posSide 填写 short ）
        side = "buy" if order['direction'] == 1 else "sell"
        params = {
            "tdMode": "cross",
            'clOrdId': str(int(time.time())),
        }
        if self.is_realtime:
            info = self.api.create_order(symbol=self.symbol, type="market", side=side, amount=order['size'], price=None, params=params)
            order1 = self.api.fetch_order(info['id'], self.symbol)
            t1 = time.time()
            while order1['status'] != "closed":
                order1 = self.api.fetch_order(info['id'], self.symbol)
                t2 = time.time()
                if (t2 - t1) > 2:
                    self.api.cancel_order(info['id'], self.symbol)  # 撤单                    
                    break

            msg = "开%s, 价格%s" % ("多" if order['direction'] == 1 else "空", order1['average'])
        else:
            msg = "开%s, 价格%s" % ("多" if order['direction'] == 1 else "空", order['in_price'])
        self.logger.info(msg)
    
    @timmer
    def sell(self, position, ticker):
        if self.is_realtime:
            # params 参数具体含义看okex.md
            sz = position['contracts']  # 默认全部
            # 平多：卖出平多（side 填写 sell；posSide 填写 long ）
            # 平空：买入平空（side 填写 buy； posSide 填写 short ）
            side = "sell" if position['side'] == "long" else "buy"
            params = {
                "tdMode": "cross",
                'clOrdId': str(int(time.time())),
            }
            self.api.create_order(symbol=self.symbol, type="market", side=side, amount=sz, price=None, params=params)
            order1 = okex.fetch_order(params['clOrdId'], self.symbol)
            msg = "平%s, 价格%s" % ("多" if position['side'] == "long" else "空", order1['average'])
        else:
            if self.deque[0]['direction'] == 1:
                r = (ticker['bid'] - self.deque[0]['in_price'] - self.deque[0]['in_price'] * 0.0005 - ticker['bid'] * 0.0005) / self.deque[0]['in_price']
            else :
                r = (self.deque[0]['in_price'] - ticker['ask'] - self.deque[0]['in_price'] * 0.0005 - ticker['bid'] * 0.0005) / self.deque[0]['in_price']
            msg = "平%s, 价格%s, 盈利%s个点" % ("多" if self.deque[0]['direction'] == 1 else "空", ticker['last'],  r)
        
        self.logger.info(msg)
        self.init()

    def hhll(self, symbol, short_len, period, logger=None):
        klines = self.get_NH_bar(symbol, period, 1)
        # df = pd.DataFrame(klines[-short_len:])
        hs = [float(line[2]) for line in klines[:short_len]]
        ls = [float(line[3]) for line in klines[:short_len]]
        cs = [float(line[4]) for line in klines[:short_len]]
        return max(hs), min(ls), hs, ls, cs

    @timmer
    def hhll_type2(self, symbol, short_len, period, logger=None):
        klines = self.get_NH_bar(symbol, period, 2)
        # df = pd.DataFrame(klines[-short_len:])
        hs = [float(line[2]) for line in klines[:short_len]]
        ls = [float(line[3]) for line in klines[:short_len]]
        cs = [float(line[4]) for line in klines[:short_len]]
        return max(hs), min(ls), hs, ls, cs 

    @timmer
    def get_ATR(self, hh, ll, hs, ls, cs):
        _atr = self.ATR(cs, hs, ls, self.short_len-1)
        # 计算ATR
        atr_up = min(_atr, hh*self.ban_out)
        atr_dw = min(_atr, ll*self.ban_out)
        return atr_up, atr_dw
    
    @timmer
    def ATR(self, close_ATR, high_ATR, low_ATR, timeperiod):
        # 计算 MTR和ATR
        close_ATR, high_ATR, low_ATR = list(close_ATR), list(high_ATR),  list(low_ATR)
        list_mtr = []
        temp = -timeperiod
        while temp < 0:
            t = max(high_ATR[temp] - low_ATR[temp], abs(close_ATR[temp - 1] - high_ATR[temp]), 
            abs(close_ATR[temp - 1] - low_ATR[temp]))
            list_mtr.append(t)
            temp += 1
        k = len(list_mtr) - timeperiod
        res = sum(list_mtr[k:]) / timeperiod
        return res