package guTa

import (
	"math"
)

const (
	UNIT_SEC   = 's'
	UNIT_MIN   = 'm'
	UNIT_HOUR  = 'h'
	UNIT_DAY   = 'd'
	UNIT_WEEK  = 'w'
	UNIT_MONTH = 'M'
	UNIT_YEAR  = 'y'
)

type TickChart struct {
	Shcode string `json:"shcode"`     // (단축)종목코드
	Ticks  []Tick `json:"price_vols"` // 가격 배열
	// Interval string `json:"interval"`   // 간격
}

type Tick struct {
	Datetime string  `json:"datetime"` // 날짜시간
	Price    float64 `json:"price"`    // 가격
	Volume   float64 `json:"volume"`   // 거래량
}

// TODO: Interval 삭제
type CandleChart struct {
	Shcode  string   `json:"shcode"`  // (단축)종목코드
	Candles []Candle `json:"candles"` // 캔들 배열
	// Interval string   `json:"interval"` // 캔들 간격
}

type Candle struct {
	Datetime string  `json:"datetime"` // 날짜시간
	Open     float64 `json:"open"`     // 시가
	High     float64 `json:"high"`     // 고가
	Low      float64 `json:"low"`      // 저가
	Close    float64 `json:"close"`    // 종가
	// Gap      float64 `json:"gap"`      // 갭등락(일봉 이상: 금일시가-전일종가, 일봉 미만: 0)
	Volume float64 `json:"volume"` // 거래량
	// Value    float64 `json:"value"`    // 거래대금
}

// *** Candle

// ** Get
func (candle *Candle) GetFields() map[string]string {
	return map[string]string{
		"Datetime": "string",
		"Open":     "float64",
		"High":     "float64",
		"Low":      "float64",
		"Close":    "float64",
		"Volume":   "float64",
	}
}

func (candle *Candle) GetDatetime() string {
	return candle.Datetime
}

func (candle *Candle) GetOpen() float64 {
	return candle.Open
}

func (candle *Candle) GetHigh() float64 {
	return candle.High
}

func (candle *Candle) GetLow() float64 {
	return candle.Low
}

func (candle *Candle) GetClose() float64 {
	return candle.Close
}

func (candle *Candle) GetVolume() float64 {
	return candle.Volume
}

// ** candle 기본 변수
// 음봉, 양봉, flat
func (candle *Candle) Sign() (sign float64) {
	sign = 0
	if candle.Close > candle.Open {
		return 1
	} else if candle.Close < candle.Open {
		return -1
	}
	return
}

// 금일 변동폭(고가-저가)
func (candle *Candle) Height() float64 {
	return candle.High - candle.Low
}

// 전일종가 대비 금일 변동폭(고가-저가) 비율
func (candle *Candle) HeightRate(gap float64) float64 {
	return (candle.High - candle.Low) / (candle.Open - gap)
}

// 몸통 길이(음수: 음봉)
func (candle *Candle) Body() (body float64) {
	body = 0
	sign := candle.Sign()
	if sign != 0 {
		body = candle.Close - candle.Open // 양수, 음수
	}
	return
}

// 몸통 비율(현재 봉 길이(키) 기준, 음수: 음봉)
func (candle *Candle) BodyRate() (bodyRate float64) {
	bodyRate = 0
	body := candle.Body()
	if body != 0 {
		bodyRate = body / (candle.High - candle.Low)
	}
	return
}

// 머리 길이(음수: 음봉)
func (candle *Candle) Head() (head float64) {
	head = 0
	sign := candle.Sign()
	if sign > 0 {
		head = candle.High - candle.Close // 양수
	} else if sign < 0 {
		head = candle.Open - candle.High // 음수
	}
	return
}

// 머리 비율(현재 봉 길이(키) 기준, 음수: 음봉)
func (candle *Candle) HeadRate() (headRate float64) {
	headRate = 0
	head := candle.Head()
	if head != 0 {
		headRate = head / (candle.High - candle.Low)
	}
	return
}

// 꼬리 길이(음수: 음봉)
func (candle *Candle) Tail() (tail float64) {
	tail = 0
	sign := candle.Sign()
	if sign > 0 {
		tail = candle.Open - candle.Low // 양수
	} else if sign < 0 {
		tail = candle.Low - candle.Close // 음수
	}
	return
}

// 꼬리 비율(현재 봉 길이(키) 기준, 음수: 음봉)
func (candle *Candle) TailRate() (tailRate float64) {
	tailRate = 0
	tail := candle.Tail()
	if tail != 0 {
		tailRate = tail / (candle.High - candle.Low)
	}
	return
}

// 전일 종가 대비 등락액
func (candle *Candle) Diff(gap float64) (diff float64) {
	return candle.Close - (candle.Open - gap) // 금종가 - 전종가
}

// 전일 종가 대비 등락률
func (candle *Candle) DiffRate(gap float64) (diffRate float64) {
	diffRate = 0
	diff := candle.Diff(gap)
	if diff != 0 {
		diffRate = diff / (candle.Open - gap)
	}
	return
}

// ** candle 변형
// * 캔들 -> 라인
func (candle *Candle) ConvToPvs() (ps []float64) {
	//
	ps = append(ps, candle.Open)
	if candle.Close > candle.Open {
		ps = append(ps, candle.Low)
		ps = append(ps, candle.High)
	} else {
		ps = append(ps, candle.High)
		ps = append(ps, candle.Low)
	}
	ps = append(ps, candle.Close)
	return
}

// *** CandleChart

// ** Get
func (candleChart *CandleChart) GetDatetimes() (datetimes []string) {
	for _, candle := range candleChart.Candles {
		datetimes = append(datetimes, candle.Datetime)
	}
	return datetimes
}

func (candleChart *CandleChart) GetOpens() (prices []float64) {
	for _, candle := range candleChart.Candles {
		prices = append(prices, candle.Open)
	}
	return prices
}

func (candleChart *CandleChart) GetHighs() (prices []float64) {
	for _, candle := range candleChart.Candles {
		prices = append(prices, candle.High)
	}
	return prices
}

func (candleChart *CandleChart) GetLows() (prices []float64) {
	for _, candle := range candleChart.Candles {
		prices = append(prices, candle.Low)
	}
	return prices
}

func (candleChart *CandleChart) GetCloses() (prices []float64) {
	for _, candle := range candleChart.Candles {
		prices = append(prices, candle.Close)
	}
	return prices
}

func (candleChart *CandleChart) GetVolumes() (volumes []float64) {
	for _, candle := range candleChart.Candles {
		volumes = append(volumes, candle.Volume)
	}
	return volumes
}

// * 캔들 압축
// 캔들을 하나로
func CandlesToOne(candles []Candle) (candle_ Candle) {
	size := len(candles)
	candle_.Open = candles[0].Open
	candle_.Close = candles[size-1].Close
	// candle_.Gap = candles[0].Gap // Gap
	sumVolume := 0.0
	// sumValue := 0.0
	highs := []float64{}
	lows := []float64{}
	for _, candle := range candles {
		highs = append(highs, candle.High)
		lows = append(lows, candle.Low)
		sumVolume += candle.Volume
		// sumValue += candle.Value
	}
	candle_.High = MaxOne(highs)
	candle_.Low = MaxOne(lows)
	candle_.Volume = sumVolume
	// candle_.Value = sumValue // Value
	// TODO: Datetime, CandleChart Interval
	return
}

// 캔들챠트 압축
// multi개 만큼의 candle을 1개로 압축
func (candleChart *CandleChart) CompressFrom(multi int) (cc CandleChart) {
	length := len(candleChart.Candles)
	if multi == 1 {
		return *candleChart
	}
	cnt := int(math.Ceil(float64(length) / float64(multi)))
	for i := 0; i < cnt; i++ {
		cc.Candles[i] = CandlesToOne(candleChart.Candles[i*multi : (i+1)*multi])
	}
	// TODO: Interval
	return
}

// 캔들챠트 압축
// to > 1 => 최종 몇개의 캔들로 ex) length: 100, to = 5: 100개 캔들 -> 5개 캔들
// to < 1 => 캔들을 비율로 줄임 ex) length: 100, to = 0.3: 100개 캔들 -> 100*0.3 = 30
func (candleChart *CandleChart) CompressTo(to float64) (cc CandleChart) {
	length := len(candleChart.Candles)
	multi := 1 // 몇 개의 캔들을 1개의 캔들로 만들것인지
	if to > 1 {
		multi = int(length / int(to)) // TODO: ceil로 변경
	} else {
		multi = int(1 / int(to))
	}

	// multi = 1이면 입력 그대로 출력
	if multi == 1 {
		return *candleChart
	}
	for i := 0; i < int(length/multi); i++ {
		cc.Candles[i] = CandlesToOne(candleChart.Candles[i*multi : (i+1)*multi])
	}
	// TODO: Interval
	return
}

// * 캔들 팽창

// ** pattern

// * 1개 candle

// * 2개 candle

// * 3개 candle
