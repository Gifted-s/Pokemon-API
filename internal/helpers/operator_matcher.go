package helpers

// URLOperatorToMongoDBOperatorMatcher returns a map that maps comparison operator from request query to mongodb operator 
func URLOperatorToMongoDBOperatorMatcher()map[string]string{
  operatorMap := map[string]string{
   "gte": "$gte",
   "lte": "$lte",
   "gt": "$gt",
   "lt": "$lt",
   "eq": "$eq",
   "ne": "$ne",
  }
  return operatorMap
}