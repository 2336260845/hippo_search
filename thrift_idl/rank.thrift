struct RankParam {
    1: required list<Doc> docs,
    2: map<string, string> extra
}

struct Doc {
    1: required string title,
    2: required string summary,
    3: required string content,
    4: required string url,
    5: double rank_score,
    6: map<string, string> extra
}

service RankService {
    list<Doc> recall(1:RankParam param)
}