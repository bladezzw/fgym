func (this *V60) OT(ticker *Ticker, Pri *WsClient, config_all *config.Config) {
	this.Ts = ticker.Ts // 注意是否有损耗
	// timetp = ts2str
	// this.Curr_date_timetuple = ""
	// if this.curr_date_timetuple > this.last_date_timetuple{
	// 	log.Printf("new date:%v\n" % str(timetp))
	// 	this.endperiod_list = this.SetPeriods(ts)
	// 	this.last_date_timetuple = this.curr_date_timetuple
	// }

	// if this.endperiod_list == []{
	// 	// 解决中断重启时没有endperiod_list的时候
	// 	this.endperiod_list = SetPeriods(ts)
	// 	this.SetPosition()
	// }

	// this.is_period_end, _ = isperiodend(ts, this.endperiod_list)

	// if not this.is_period_end && this.last_period_end: {
	// 	//# 前一秒在范围里面,这一秒不在范围里面,则满足barhhll更新条件
	// 	// # 更新最近short_len个period bar的最高点,最低点
	// 	this.floor_h, this.floor_l, hs, ls, cs = this.hhll(this.symbol, this.short_len, this.period, this.logger)
	// 	// # 计算ATR
	// 	atr_up, atr_dw = this.Set_ATR(this.floor_h, this.floor_l, hs, ls, cs)
	// }

	// // ######################################  多  ######################################
	// if (this.position.get("side", 0) == "long"){
	// 	if this.ups.out_price >= ticker.Bid {
	// 		// # 买一价小于等于出场价格
	// 		log.Printf("%v, 多头出场, 目标价位%v, 实际价位%v", this.symbol, this.ups.out_price, ticker.BidPx)
	// 		this.ups.sell(this.position, ticker)
	// 		this.get_position_usd( this.symbol)
	// 		this.last_up_out_minute = ts2tp(ts).strftime("%Y-%m-%d %H:%M")
	// 	}else{
	// 		if this.ups.hh < ticker.Last{
	// 			this.ups.hh = ticker.Last
	// 			this.ups.out_price = this.ups.hh - this.ups.max_atr
	// 		}
	// }
	// if this.position.get("side", 0) != "long" {
	// 	if this.floor_h <= ticker.AskPx{
	// 		// # 卖一价大于或等于之前bar的最高点
	// 		if ts2tp(ts).strftime("%Y-%m-%d %H:%M") != this.last_up_out_minute{
	// 			this.logger.info("%s, 多头进场, 目标价位%s, 实际价位%s"% (this.symbol, this.floor_h, ticker.AskPx))
	// 			size = round(this.Value / (this.symbol_info.ContractSize * ticker.AskPx))
	// 			UP = dict(direction=1, dt=ts, in_price=ticker.AskPx, size=size)
	// 			this.ups.append(UP)
	// 			this.ups.hh = ticker.AskPx
	// 			this.ups.max_atr = max(this.ups.max_atr, atr_up)
	// 			this.ups.out_price = this.ups.hh - this.ups.max_atr
	// 		}
	// 	}
	// }
	// // ######################################  short  ######################################
	// if this.position.get("side", 0) == "short"{
	// //  # shortout
	// 	if this.dws.out_price <= ticker.AskPx{
	// 		// # 卖一价大于等于出场价格
	// 		this.logger.info("%s, 空头出场, 目标价位%s, 实际价位%s"% (this.symbol, this.dws.out_price, ticker.AskPx))
	// 		this.dws.sell(position, ticker)
	// 		this.get_position_usd(this.symbol)
	// 		this.last_dw_out_minute = ts2tp(ts).strftime("%Y-%m-%d %H:%M")
	// 	}else{
	// 		if this.dws.ll > ticker.Last{
	// 			this.dws.ll = ticker.Last
	// 			this.dws.out_price = this.dws.ll + this.dws.max_atr
	// 		}
	// 	}
	// }
	// if this.position.get("side", 0) != "short"{
	// 	if this.floor_l >= ticker.BidPx{
	// 		//# 买一价小于或等于之前bar的最低点:
	// 		if ts2tp(ts).strftime("%Y-%m-%d %H:%M") != this.last_dw_out_minute{
	// 			log.Printf("%v, 空头进场, 目标价位%v, 实际价位%v\n", this.symbol, this.floor_l, ticker.BidPx)
	// 			size = round(this.Value / (this.symbol_info.ContractSize * ticker.BidPx))
	// 			DW = dict(direction=-1,dt=ts, in_price=ticker.BidPx, size=size)
	// 			this.dws.append(DW)
	// 			this.dws.ll = ticker.BidPx
	// 			this.dws.max_atr = max(this.dws.max_atr, atr_dw)
	// 			this.dws.out_price = this.dws.ll + this.dws.max_atr
	// 		}
	// 	}
	// }
	// // ############################################
	if this.Floor_h < ticker.Last {
		// # update hh
		this.Floor_h = ticker.Last
	}
	if this.Floor_l > ticker.Last {
		this.Floor_l = ticker.Last
	}
	// this.Last_period_end = Is_period_end
}
