package utils


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