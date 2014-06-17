package tools

type Exception struct {
	Server string
	Msg    string
}

var Exceptions = map[string]Exception{
	"al":               Exception{"www.akep.al", "web (CAPTCHA) http://www.akep.al/sq/kerkoni-domain"},
	"ar":               Exception{"nic.ar", "web (CAPTCHA)"},
	"az":               Exception{"www.whois.az", "web (POST)"},
	"ba":               Exception{"nic.ba", "web (POST)"},
	"bd":               Exception{"www.whois.com.bd", "web"},
	"bm":               Exception{"www.bermudanic.bm", "web (POST)"},
	"bs":               Exception{"www.register.bs", "web"},
	"bt":               Exception{"www.nic.bt", "web (POST)"},
	"bv":               Exception{"whois.norid.no", "http://www.norid.no/navnepolitikk.en.html#link1"},
	"cm":               Exception{"antic.cm", "web (sessioned)"},
	"cr":               Exception{"www.nic.cr", "web (CAPTCHA)"},
	"cu":               Exception{"www.nic.cu", "web"},
	"cv":               Exception{"www.dns.cv", "web"},
	"cy":               Exception{"www.nic.cy", "web"},
	"dj":               Exception{"www.dj", "web"},
	"eg":               Exception{"lookup.egregistry.eg", "web (POST)"},
	"fj":               Exception{"domains.fj", "web "},
	"fk":               Exception{"whois.marcaria.com", "web"},
	"ge":               Exception{"www.nic.net.ge", "web (POST)"},
	"gf":               Exception{"www.dom-enic.com", "web (POST)"},
	"gm":               Exception{"www.nic.gm", "web"},
	"gmo":              Exception{"whois.gmoregistry.net", "http://en.wikipedia.org/wiki/.gmo"},
	"gp":               Exception{"www.dom-enic.com", "web (POST)"},
	"gr":               Exception{"grweb.ics.forth.gr", "web (CAPTCHA)"},
	"gt":               Exception{"www.gt", "web"},
	"kw":               Exception{"www.kw", "web (POST)"},
	"ky":               Exception{"kynseweb.messagesecure.com", "web (POST)"},
	"lb":               Exception{"www.aub.edu.lb", "web"},
	"lc":               Exception{"www.nic.lc", "web (POST)"},
	"lk":               Exception{"whois.nic.lk", "http://nic.lk"},
	"ls":               Exception{"www.co.ls", "web"},
	"mq":               Exception{"www.dom-enic.com", "web (POST)"},
	"mt":               Exception{"www.nic.org.mt", "web"},
	"mw":               Exception{"www.registrar.mw", "web"},
	"mz":               Exception{"www.domains.co.mz", "web (CAPTCHA)"},
	"ni":               Exception{"www.nic.ni", "web (POST AJAX)"},
	"np":               Exception{"register.mos.com.np", "web (POST)"},
	"nr":               Exception{"cenpac.net.nr", "http://cenpac.net.nr/dns/"},
	"pa":               Exception{"www.nic.pa", "web"},
	"ph":               Exception{"www.dot.ph", "web (CAPTCHA)"},
	"pk":               Exception{"pk6.pknic.net.pk", "web"},
	"pn":               Exception{"www.government.pn", "web (POST)"},
	"py":               Exception{"www.nic.py", "web (POST)"},
	"rw":               Exception{"whois.ricta.org.rw", "web"},
	"sj":               Exception{"whois.norid.no", "http://www.norid.no/navnepolitikk.en.html#link1"},
	"sl":               Exception{"www.nic.sl", "web"},
	"sr":               Exception{"www.register.sr", "web (CAPTCHA)"},
	"tg":               Exception{"www.netmaster.tg", "web"},
	"tj":               Exception{"www.nic.tj", "web"},
	"tt":               Exception{"www.nic.tt", "web"},
	"va":               Exception{"whois.iana.org", "Every .va domain name owned by the Vatican"},
	"vi":               Exception{"secure.nic.vi", "web (POST)"},
	"vn":               Exception{"whois.vnnic.vn", "web"},
	"xn--90a3ac":       Exception{"whois.rnids.rs", "http://en.wikipedia.org/wiki/.xn--90a3ac"},
	"xn--fzc2c9e2c":    Exception{"whois.nic.lk", "http://www.iana.org/domains/root/db/.xn--fzc2c9e2c.html"},
	"xn--mgbc0a9azcg":  Exception{"whois.iam.net.ma", "Morocco"},
	"xn--pgbs0dh":      Exception{"whois.ati.tn", "Tunisia"},
	"xn--rhqv96g":      Exception{"whois.nic.xn--rhqv96g", "China"},
	"xn--ses554g":      Exception{"whois.gtld.knet.cn", "China: whois -h whois.gtld.knet.cn nic.xn--ses554g"},
	"xn--xkc2al3hye2a": Exception{"whois.nic.lk", "Sri Lanka"},
	"za":               Exception{"whois.registry.net.za", "http://en.wikipedia.org/wiki/.za"},
	"zw":               Exception{"www.zispa.org.zw", "web"},

	// From Ruby Whois
	// https://github.com/weppos/whois/blob/master/data/tld.json
	"ae.org":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"ar.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"br.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"cn.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"com.de":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"de.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"eu.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"gb.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"gb.net":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"gr.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"hu.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"jp.net":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"jpn.com":                  Exception{"whois.centralnic.com", "Ruby Whois"},
	"kr.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"no.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"qc.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"ru.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"sa.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"se.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"se.net":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"uk.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"uk.net":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"us.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"us.org":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"uy.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"za.com":                   Exception{"whois.centralnic.com", "Ruby Whois"},
	"za.net":                   Exception{"whois.za.net", "Ruby Whois"},
	"eu.org":                   Exception{"whois.eu.org", "Ruby Whois"},
	"za.org":                   Exception{"whois.za.org", "Ruby Whois"},
	"e164.arpa":                Exception{"whois.ripe.net", "Ruby Whois"},
	"in-addr.arpa":             Exception{"", "Ruby Whois"},
	"priv.at":                  Exception{"whois.nic.priv.at", "Ruby Whois"},
	"co.ca":                    Exception{"whois.co.ca", "Ruby Whois"},
	"edu.cn":                   Exception{"whois.edu.cn", "Ruby Whois"},
	"aeroport.fr":              Exception{"whois.smallregistry.net", "Ruby Whois"},
	"avocat.fr":                Exception{"whois.smallregistry.net", "Ruby Whois"},
	"chambagri.fr":             Exception{"whois.smallregistry.net", "Ruby Whois"},
	"chirurgiens-dentistes.fr": Exception{"whois.smallregistry.net", "Ruby Whois"},
	"experts-comptables.fr":    Exception{"whois.smallregistry.net", "Ruby Whois"},
	"geometre-expert.fr":       Exception{"whois.smallregistry.net", "Ruby Whois"},
	"medecin.fr":               Exception{"whois.smallregistry.net", "Ruby Whois"},
	"notaires.fr":              Exception{"whois.smallregistry.net", "Ruby Whois"},
	"pharmacien.fr":            Exception{"whois.smallregistry.net", "Ruby Whois"},
	"port.fr":                  Exception{"whois.smallregistry.net", "Ruby Whois"},
	"veterinaire.fr":           Exception{"whois.smallregistry.net", "Ruby Whois"},
	"co.pl":                    Exception{"whois.co.pl", "Ruby Whois"},
	"edu.ru":                   Exception{"whois.informika.ru", "Ruby Whois"},
	"in.ua":                    Exception{"whois.in.ua", "Ruby Whois"},
	"ac.uk":                    Exception{"whois.ja.net", "Ruby Whois"},
	"bl.uk":                    Exception{"", "Ruby Whois"},
	"british-library.uk":       Exception{"", "Ruby Whois"},
	"gov.uk":                   Exception{"whois.ja.net", "Ruby Whois"},
	"icnet.uk":                 Exception{"", "Ruby Whois"},
	"jet.uk":                   Exception{"", "Ruby Whois"},
	"mod.uk":                   Exception{"", "Ruby Whois"},
	"nhs.uk":                   Exception{"", "Ruby Whois"},
	"nls.uk":                   Exception{"", "Ruby Whois"},
	"parliament.uk":            Exception{"", "Ruby Whois"},
	"police.uk":                Exception{"", "Ruby Whois"},
	"com.uy":                   Exception{"nic.anteldata.com.uy", "https://nic.anteldata.com.uy/dns/consultaWhois/whois.action"},
	"ac.za":                    Exception{"whois.ac.za", "Ruby Whois"},
	"co.za":                    Exception{"whois.registry.net.za", "Ruby Whois"},
	"gov.za":                   Exception{"whois.gov.za", "Ruby Whois"},
	"org.za":                   Exception{"whois.org.za", "Ruby Whois"},
}
