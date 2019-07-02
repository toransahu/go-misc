package ch1

type Apollo interface {
    About() string
}

type PAPI struct {
    apollo Apollo
}

func (papi *PAPI) GetApolloVersion() string  {
    return papi.apollo.About()
}
