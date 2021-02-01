struct RecallParam {
    1: required string query,
    2: map<string, string> extra
}

struct Doc {
    1: required string title,
    2: required string summary,
    3: required string content,
    4: required string url,
    5: map<string, string> extra
}

service RecallService {
    list<Doc> recall(1:RecallParam param)
}