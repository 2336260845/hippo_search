struct QueryParam {
    1: required string query,
    2: required string analysis,
    3: map<string, string> extra
}

service QueryAnalysisService {
    list<string> queryAnalysis(1:QueryParam param)
}