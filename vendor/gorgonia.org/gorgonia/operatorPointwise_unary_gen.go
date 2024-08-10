package gorgonia

// Code generated by genapi, which is a API generation tool for Gorgonia. DO NOT EDIT.

func (f *sf32UnaryOperator) unaryOpType() ʘUnaryOperatorType {

	switch f {
	case &absf32:
		return absOpType
	case &signf32:
		return signOpType
	case &ceilf32:
		return ceilOpType
	case &floorf32:
		return floorOpType
	case &sinf32:
		return sinOpType
	case &cosf32:
		return cosOpType
	case &expf32:
		return expOpType
	case &lnf32:
		return lnOpType
	case &log2f32:
		return log2OpType
	case &negf32:
		return negOpType
	case &squaref32:
		return squareOpType
	case &sqrtf32:
		return sqrtOpType
	case &inversef32:
		return inverseOpType
	case &inverseSqrtf32:
		return inverseSqrtOpType
	case &cubef32:
		return cubeOpType
	case &tanhf32:
		return tanhOpType
	case &sigmoidf32:
		return sigmoidOpType
	case &log1pf32:
		return log1pOpType
	case &expm1f32:
		return expm1OpType
	case &softplusf32:
		return softplusOpType
	}
	return maxʘUnaryOperator
}
func (f *sf32UnaryOperator) String() string { return f.unaryOpType().String() }

func (f *sf64UnaryOperator) unaryOpType() ʘUnaryOperatorType {

	switch f {
	case &absf64:
		return absOpType
	case &signf64:
		return signOpType
	case &ceilf64:
		return ceilOpType
	case &floorf64:
		return floorOpType
	case &sinf64:
		return sinOpType
	case &cosf64:
		return cosOpType
	case &expf64:
		return expOpType
	case &lnf64:
		return lnOpType
	case &log2f64:
		return log2OpType
	case &negf64:
		return negOpType
	case &squaref64:
		return squareOpType
	case &sqrtf64:
		return sqrtOpType
	case &inversef64:
		return inverseOpType
	case &inverseSqrtf64:
		return inverseSqrtOpType
	case &cubef64:
		return cubeOpType
	case &tanhf64:
		return tanhOpType
	case &sigmoidf64:
		return sigmoidOpType
	case &log1pf64:
		return log1pOpType
	case &expm1f64:
		return expm1OpType
	case &softplusf64:
		return softplusOpType
	}
	return maxʘUnaryOperator
}
func (f *sf64UnaryOperator) String() string { return f.unaryOpType().String() }
