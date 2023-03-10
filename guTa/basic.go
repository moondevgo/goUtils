// C:\MoonDev\withLang\inGo\goProject\gpStock\_refs\mp_stock\__analysis_.py
// C:\MoonDev\withLang\inGo\goProject\gpStock\_refs\mp_stock\analysis_technicals.py

package guTa

import (
	"log"
	"math"
)

type Integer interface {
	int | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

// ** basic Math
// 소수점 자리수(precision) 이하 제거
// precision: 양수: 소수점 이하 몇 번째 자리, 음수: 소수점 이상 몇 번째 자리
func FixFloat(num float64, precision int, roundType byte) float64 {
	var rounded float64
	output := math.Pow(10, float64(precision))
	switch roundType {
	case 'r': // 반올림
		rounded = math.Round(num * output)
	case 'c': // 올림
		rounded = math.Ceil(num * output)
	case 'f': // 버림
		rounded = math.Floor(num * output)
	}

	return rounded / output
}

func Round_(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}

func Ceil_(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Ceil(num*output)) / output
}

func Floor_(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Floor(num*output)) / output
}

// ** 속도.가속도
// 속도(평균변화율)
func Delta1[T Number](data []T) (data_ []T) {
	data_ = []T{0}
	for i, d := range data[1:] {
		data_ = append(data_, d-data[i])
	}
	return
}

// 가속도(2계도함수)
func Delta2[T Number](data []T) (data_ []T) {
	data_ = Delta1(Delta1(data))
	data_[1] = 0
	return data_
}

// 속도(평균변화율)
func Delta1_[T Number](data []T, precision int) (data_ []T) {
	data_ = []T{0}
	// func Delta1_[T Number](data []T, precision int) (data_ []float64) {
	// 	data_ = []float64{0}
	for i, d := range data[1:] {
		data_ = append(data_, T(Round_(float64(d-data[i]), precision)))
	}
	return
}

// 가속도(2계도함수)
func Delta2_[T Number](data []T, precision int) (data_ []T) {
	data_ = Delta1_(Delta1_(data, precision), precision)
	data_[1] = 0
	return data_
}

// ** 최대.최소값
// * 최대값
// 최대값(인덱스, 최대값)
func MaxOne_(values []float64) (idx int, max float64) {
	max = math.Inf(-1)
	idx = 0
	for i, value := range values {
		// TODO: 최대값 최초(value > max), 최후 인덱스 (value >= max)
		if value > max {
			max = value
			idx = i
		}
	}
	return
}

// 최대값
func MaxOne(values []float64) (max float64) {
	max = math.Inf(-1)
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return
}

// 최대값 전체(인덱스, 최대값 배열)
func MaxAll_(values []float64) (idxs []int, max float64) {
	max = math.Inf(-1)
	idxs = []int{0}
	for i, value := range values {
		if value > max {
			log.Printf("\n### MaxAll First Max: %v\n", value)
			max = value
			idxs = []int{i}
		} else if value == max {
			log.Printf("\n**** MaxAll Multi Max: %v\n", value)
			idxs = append(idxs, i)
		}
	}
	return
}

// * 최소값
// 최소값(인덱스, 최대값)
func MinOne_(values []float64) (idx int, min float64) {
	min = math.Inf(1)
	idx = 0
	for i, value := range values {
		if value < min {
			min = value
			idx = i
		}
	}
	return
}

// 최소값
func MinOne(values []float64) (min float64) {
	min = math.Inf(1)
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return
}

// 최소값 전체(인덱스, 최소값 배열)
func MinAll_(values []float64) (idxs []int, min float64) {
	min = math.Inf(1)
	idxs = []int{0}
	for i, value := range values {
		if value < min {
			min = value
			idxs = []int{i}
		} else if value == min {
			idxs = append(idxs, i)
		}
	}
	return
}

// ** 극대.극소값
// * 극소값
func ExtremeMins(xs, ys []float64) (xs_, ys_ []float64) {
	size := len(ys)
	if ys[1] >= ys[0] {
		xs_ = append(xs_, xs[0])
		ys_ = append(ys_, ys[0])
	}
	for i := 1; i < size-1; i++ {
		if ys[i] <= ys[i-1] && ys[i] <= ys[i+1] {
			xs_ = append(xs_, xs[i])
			ys_ = append(ys_, ys[i])
		}
	}
	if ys[size-1] <= ys[size-2] {
		xs_ = append(xs_, xs[size-1])
		ys_ = append(ys_, ys[size-1])
	}

	return
}

// * 극대값
func ExtremeMaxs(xs, ys []float64) (xs_, ys_ []float64) {
	size := len(ys)
	if ys[1] < ys[0] {
		xs_ = append(xs_, xs[0])
		ys_ = append(ys_, ys[0])
	}
	for i := 1; i < size-1; i++ {
		if ys[i] >= ys[i-1] && ys[i] >= ys[i+1] {
			xs_ = append(xs_, xs[i])
			ys_ = append(ys_, ys[i])
		}
	}
	if ys[size-1] >= ys[size-2] {
		xs_ = append(xs_, xs[size-1])
		ys_ = append(ys_, ys[size-1])
	}
	return
}

// ** 추세선
// 2점 -> 기울기
func Gradient(x0, y0, x1, y1 float64) (m float64) {
	return (y1 - y0) / (x1 - x0)
}

// 2점 -> y절편
func InterceptY(x0, y0, x1, y1 float64) (b float64) {
	return y0 - (y1-y0)*x0/(x1-x0)
}

// 2점 -> 기울기, y절편
func GradientInterceptY(x0, y0, x1, y1 float64) (m, b float64) {
	m = (y1 - y0) / (x1 - x0)
	return m, y0 - m*x0
}

// x값, 기울기, y절편
func SlopeLine(indexes []float64, m, b float64) (ys_ []float64) {
	for _, idx := range indexes {
		ys_ = append(ys_, m*idx+b)
	}
	return
}

// * 추세선(극소값)
func TrenlineMins(xs, ys []float64) {

}

// * Cross 저항.지지선 돌파
func Cross() {

}

// ** 통계 관련
// 평균
// type Number interface {
//     constraints.Float | constraints.Integer
// }

// func mean[T Number](data []T) float64 {
//     if len(data) == 0 {
//         return 0
//     }
//     var sum float64
//     for _, d := range data {
//         sum += float64(d)
//     }
//     return sum / float64(len(data))
// }

func Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}

// 분산
func VAR(data []float64) (var_ float64) {
	size := len(data)
	if size == 0 {
		return 0
	}
	mean := Mean(data)
	for _, d := range data {
		var_ += math.Pow(d-mean, 2)
	}
	return var_ / float64(size)
}

// 표준편차
func STD(data []float64) (std_ float64) {
	size := len(data)
	if size == 0 {
		return 0
	}

	return math.Sqrt(VAR(data))
}

// * 정규화
// 정규화 정규분포
func NormNormal(data []float64) (data_ []float64) {
	mean := Mean(data)
	std := STD(data)
	if std == 0 {
		for i := 0; i < len(data); i++ {
			data_ = append(data_, 0)
		}
		return data_
	}
	for _, d := range data {
		data_ = append(data_, (d-mean)/std)
	}
	return
}

// 가격 -> (0, 1)
func NormMaxMin(data []float64) (data_ []float64) {
	max := MaxOne(data)
	min := MinOne(data)
	if max == min {
		for i := 0; i < len(data); i++ {
			data_ = append(data_, 0)
		}
		return data_
	}
	for _, d := range data {
		data_ = append(data_, (d-min)/(max-min))
	}
	return
}

// 기준가 대비 차액(률)
func NormSimple(data []float64, index int) (data_ []float64) {
	base := data[index]
	for _, d := range data {
		data_ = append(data_, (d-base)/base)
	}
	return
}

// def _norm_nd(series):
//     """정규화 by 정규분포
//     """
//     return (series - series.mean())/series.std()

// def _norm_mm(series):
//     """정규화 by 최대.최소
//     """
//     return (series - series.min())/(series.max() - series.min())

// def _norm(df, norms={"Open": "nd", "Close": "nd", "Volume": "mm"}):
//     for column, norm in norms.items():
//         if norm == "nd":
//             df[column] = _norm_nd(df[column])
//         else:
//             df[column] = _norm_mm(df[column])
//     return df

// def _amplitude(df):
//     df['Amp'] = (df['High'] - df['Low'])/(df['High'] + df['Low'])
//     return df

// * 상관계수
func CorrelationCoefficient(X []int, Y []int, n int) float64 {

	sum_X := 0
	sum_Y := 0
	sum_XY := 0
	squareSum_X := 0
	squareSum_Y := 0

	for i := 0; i < n; i++ {
		// sum of elements of array X.
		sum_X = sum_X + X[i]

		// sum of elements of array Y.
		sum_Y = sum_Y + Y[i]

		// sum of X[i] * Y[i].
		sum_XY = sum_XY + X[i]*Y[i]

		// sum of square of array elements.
		squareSum_X = squareSum_X + X[i]*X[i]
		squareSum_Y = squareSum_Y + Y[i]*Y[i]
	}

	// use formula for calculating correlation
	// coefficient.
	corr := float64((n*sum_XY - sum_X*sum_Y)) /
		(math.Sqrt(float64((n*squareSum_X - sum_X*sum_X) * (n*squareSum_Y - sum_Y*sum_Y))))

	return corr

}

// ** 회계 관련
// * Maximum DrawDown
func FindMdd(values []float64) (val float64) {
	val = 1
	size := len(values)
	for i := 1; i < size; i++ {
		bw_max := MaxOne(values[i:])
		curr := values[i]
		mdd := curr / bw_max

		if mdd < val {
			val = mdd
		}
	}
	return
}

// def get_mdd(x):
//     """
//     MDD(Maximum Draw-Down)
//     :return: (peak_upper, peak_lower, mdd rate)
//     """
//     arr_v = np.array(x)
//     peak_lower = np.argmax(np.maximum.accumulate(arr_v) - arr_v)
//     peak_upper = np.argmax(arr_v[:peak_lower])
//     return (peak_upper, peak_lower, (arr_v[peak_lower] - arr_v[peak_upper]) / arr_v[peak_upper])

// def cum_mean(arr):
//     cum_sum = np.cumsum(arr, axis=0)
//     for i in range(cum_sum.shape[0]):
//         if i == 0:
//             continue
//         print(cum_sum[i] / (i + 1))
//         cum_sum[i] =  cum_sum[i] / (i + 1)
//     return cum_sum
