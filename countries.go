package gox

import (
	"github.com/gopub/log"
	"sync"
)

type Country struct {
	Name        string `json:"name"`
	CallingCode int    `json:"calling_code"`
	Flag        string `json:"flag"`
}

func GetCountries() []*Country {
	loadCountries()
	return countries
}

func GetCountryByCallingCode(code int) *Country {
	loadCountries()
	return codeToCountry[code]
}

var countries []*Country
var codeToCountry map[int]*Country
var loadCountriesMutex sync.Mutex

func loadCountries() {
	if countries != nil {
		return
	}

	loadCountriesMutex.Lock()
	defer loadCountriesMutex.Unlock()

	if countries != nil {
		return
	}

	if err := JSONUnmarshal([]byte(countriesJSONString), &countries); err != nil {
		log.Error(err)
		countries = []*Country{}
		return
	}

	codeToCountry = make(map[int]*Country, len(countries))
	for _, c := range countries {
		codeToCountry[c.CallingCode] = c
	}
}

var countriesJSONString = `
[
	{
		"name": "Afghanistan",
		"flag": "🇦🇫",
		"calling_code": 93
	},
	{
		"name": "Albania",
		"flag": "🇦🇱",
		"calling_code": 355
	},
	{
		"name": "Algeria",
		"flag": "🇩🇿",
		"calling_code": 213
	},
	{
		"name": "American Samoa",
		"flag": "🇦🇸",
		"calling_code": 1684
	},
	{
		"name": "Andorra",
		"flag": "🇦🇩",
		"calling_code": 376
	},
	{
		"name": "Angola",
		"flag": "🇦🇴",
		"calling_code": 244
	},
	{
		"name": "Anguilla",
		"flag": "🇦🇮",
		"calling_code": 1264
	},
	{
		"name": "Antigua and Barbuda",
		"flag": "🇦🇬",
		"calling_code": 1268
	},
	{
		"name": "Argentina",
		"flag": "🇦🇷",
		"calling_code": 54
	},
	{
		"name": "Armenia",
		"flag": "🇦🇲",
		"calling_code": 374
	},
	{
		"name": "Aruba",
		"flag": "🇦🇼",
		"calling_code": 297
	},
	{
		"name": "Australia",
		"flag": "🇦🇺",
		"calling_code": 61
	},
	{
		"name": "Austria",
		"flag": "🇦🇹",
		"calling_code": 43
	},
	{
		"name": "Azerbaijan",
		"flag": "🇦🇿",
		"calling_code": 994
	},
	{
		"name": "Bahamas",
		"flag": "🇧🇸",
		"calling_code": 1242
	},
	{
		"name": "Bahrain",
		"flag": "🇧🇭",
		"calling_code": 973
	},
	{
		"name": "Bangladesh",
		"flag": "🇧🇩",
		"calling_code": 880
	},
	{
		"name": "Barbados",
		"flag": "🇧🇧",
		"calling_code": 1246
	},
	{
		"name": "Belarus",
		"flag": "🇧🇾",
		"calling_code": 375
	},
	{
		"name": "Belgium",
		"flag": "🇧🇪",
		"calling_code": 32
	},
	{
		"name": "Belize",
		"flag": "🇧🇿",
		"calling_code": 501
	},
	{
		"name": "Benin",
		"flag": "🇧🇯",
		"calling_code": 229
	},
	{
		"name": "Bermuda",
		"flag": "🇧🇲",
		"calling_code": 1441
	},
	{
		"name": "Bhutan",
		"flag": "🇧🇹",
		"calling_code": 975
	},
	{
		"name": "Bolivia",
		"flag": "🇧🇴",
		"calling_code": 591
	},
	{
		"name": "Botswana",
		"flag": "🇧🇼",
		"calling_code": 267
	},
	{
		"name": "Brazil",
		"flag": "🇧🇷",
		"calling_code": 55
	},
	{
		"name": "British Indian Ocean Territory",
		"flag": "🇮🇴",
		"calling_code": 246
	},
	{
		"name": "British Virgin Islands",
		"flag": "🇻🇬",
		"calling_code": 1284
	},
	{
		"name": "Brunei",
		"flag": "🇧🇳",
		"calling_code": 673
	},
	{
		"name": "Bulgaria",
		"flag": "🇧🇬",
		"calling_code": 359
	},
	{
		"name": "Burkina Faso",
		"flag": "🇧🇫",
		"calling_code": 226
	},
	{
		"name": "Burundi",
		"flag": "🇧🇮",
		"calling_code": 257
	},
	{
		"name": "Cambodia",
		"flag": "🇰🇭",
		"calling_code": 855
	},
	{
		"name": "Cameroon",
		"flag": "🇨🇲",
		"calling_code": 237
	},
	{
		"name": "Canada",
		"flag": "🇨🇦",
		"calling_code": 1
	},
	{
		"name": "Cape Verde",
		"flag": "🇨🇻",
		"calling_code": 238
	},
	{
		"name": "Cayman Islands",
		"flag": "🇰🇾",
		"calling_code": 1345
	},
	{
		"name": "Central African Republic",
		"flag": "🇨🇫",
		"calling_code": 236
	},
	{
		"name": "Chad",
		"flag": "🇹🇩",
		"calling_code": 235
	},
	{
		"name": "Chile",
		"flag": "🇨🇱",
		"calling_code": 56
	},
	{
		"name": "China",
		"flag": "🇨🇳",
		"calling_code": 86
	},
	{
		"name": "Christmas Island",
		"flag": "🇨🇽",
		"calling_code": 61
	},
	{
		"name": "Colombia",
		"flag": "🇨🇴",
		"calling_code": 57
	},
	{
		"name": "Comoros",
		"flag": "🇰🇲",
		"calling_code": 269
	},
	{
		"name": "Cook Islands",
		"flag": "🇨🇰",
		"calling_code": 682
	},
	{
		"name": "Costa Rica",
		"flag": "🇨🇷",
		"calling_code": 506
	},
	{
		"name": "Croatia",
		"flag": "🇭🇷",
		"calling_code": 385
	},
	{
		"name": "Cuba",
		"flag": "🇨🇺",
		"calling_code": 53
	},
	{
		"name": "Cyprus",
		"flag": "🇨🇾",
		"calling_code": 357
	},
	{
		"name": "Czech Republic",
		"flag": "🇨🇿",
		"calling_code": 420
	},
	{
		"name": "Denmark",
		"flag": "🇩🇰",
		"calling_code": 45
	},
	{
		"name": "Djibouti",
		"flag": "🇩🇯",
		"calling_code": 253
	},
	{
		"name": "Dominica",
		"flag": "🇩🇲",
		"calling_code": 1767
	},
	{
		"name": "Ecuador",
		"flag": "🇪🇨",
		"calling_code": 593
	},
	{
		"name": "Egypt",
		"flag": "🇪🇬",
		"calling_code": 20
	},
	{
		"name": "El Salvador",
		"flag": "🇸🇻",
		"calling_code": 503
	},
	{
		"name": "Equatorial Guinea",
		"flag": "🇬🇶",
		"calling_code": 240
	},
	{
		"name": "Eritrea",
		"flag": "🇪🇷",
		"calling_code": 291
	},
	{
		"name": "Estonia",
		"flag": "🇪🇪",
		"calling_code": 372
	},
	{
		"name": "Ethiopia",
		"flag": "🇪🇹",
		"calling_code": 251
	},
	{
		"name": "Falkland Islands",
		"flag": "🇫🇰",
		"calling_code": 500
	},
	{
		"name": "Faroe Islands",
		"flag": "🇫🇴",
		"calling_code": 298
	},
	{
		"name": "Fiji",
		"flag": "🇫🇯",
		"calling_code": 679
	},
	{
		"name": "Finland",
		"flag": "🇫🇮",
		"calling_code": 358
	},
	{
		"name": "France",
		"flag": "🇫🇷",
		"calling_code": 33
	},
	{
		"name": "French Guiana",
		"flag": "🇬🇫",
		"calling_code": 594
	},
	{
		"name": "French Polynesia",
		"flag": "🇵🇫",
		"calling_code": 689
	},
	{
		"name": "Gabon",
		"flag": "🇬🇦",
		"calling_code": 241
	},
	{
		"name": "Gambia",
		"flag": "🇬🇲",
		"calling_code": 220
	},
	{
		"name": "Georgia",
		"flag": "🇬🇪",
		"calling_code": 995
	},
	{
		"name": "Germany",
		"flag": "🇩🇪",
		"calling_code": 49
	},
	{
		"name": "Ghana",
		"flag": "🇬🇭",
		"calling_code": 233
	},
	{
		"name": "Gibraltar",
		"flag": "🇬🇮",
		"calling_code": 350
	},
	{
		"name": "Greece",
		"flag": "🇬🇷",
		"calling_code": 30
	},
	{
		"name": "Greenland",
		"flag": "🇬🇱",
		"calling_code": 299
	},
	{
		"name": "Grenada",
		"flag": "🇬🇩",
		"calling_code": 1473
	},
	{
		"name": "Guadeloupe",
		"flag": "🇬🇵",
		"calling_code": 590
	},
	{
		"name": "Guam",
		"flag": "🇬🇺",
		"calling_code": 1671
	},
	{
		"name": "Guatemala",
		"flag": "🇬🇹",
		"calling_code": 502
	},
	{
		"name": "Guernsey",
		"flag": "🇬🇬",
		"calling_code": 44
	},
	{
		"name": "Guinea",
		"flag": "🇬🇳",
		"calling_code": 224
	},
	{
		"name": "Guinea-Bissau",
		"flag": "🇬🇼",
		"calling_code": 245
	},
	{
		"name": "Guyana",
		"flag": "🇬🇾",
		"calling_code": 592
	},
	{
		"name": "Haiti",
		"flag": "🇭🇹",
		"calling_code": 509
	},
	{
		"name": "Honduras",
		"flag": "🇭🇳",
		"calling_code": 504
	},
	{
		"name": "Hong Kong",
		"flag": "🇭🇰",
		"calling_code": 852
	},
	{
		"name": "Hungary",
		"flag": "🇭🇺",
		"calling_code": 36
	},
	{
		"name": "Iceland",
		"flag": "🇮🇸",
		"calling_code": 354
	},
	{
		"name": "India",
		"flag": "🇮🇳",
		"calling_code": 91
	},
	{
		"name": "Indonesia",
		"flag": "🇮🇩",
		"calling_code": 62
	},
	{
		"name": "Iran",
		"flag": "🇮🇷",
		"calling_code": 98
	},
	{
		"name": "Iraq",
		"flag": "🇮🇶",
		"calling_code": 964
	},
	{
		"name": "Ireland",
		"flag": "🇮🇪",
		"calling_code": 353
	},
	{
		"name": "Isle of Man",
		"flag": "🇮🇲",
		"calling_code": 44
	},
	{
		"name": "Israel",
		"flag": "🇮🇱",
		"calling_code": 972
	},
	{
		"name": "Italy",
		"flag": "🇮🇹",
		"calling_code": 39
	},
	{
		"name": "Jamaica",
		"flag": "🇯🇲",
		"calling_code": 1876
	},
	{
		"name": "Japan",
		"flag": "🇯🇵",
		"calling_code": 81
	},
	{
		"name": "Jersey",
		"flag": "🇯🇪",
		"calling_code": 44
	},
	{
		"name": "Jordan",
		"flag": "🇯🇴",
		"calling_code": 962
	},
	{
		"name": "Kenya",
		"flag": "🇰🇪",
		"calling_code": 254
	},
	{
		"name": "Kiribati",
		"flag": "🇰🇮",
		"calling_code": 686
	},
	{
		"name": "Kuwait",
		"flag": "🇰🇼",
		"calling_code": 965
	},
	{
		"name": "Kyrgyzstan",
		"flag": "🇰🇬",
		"calling_code": 996
	},
	{
		"name": "Laos",
		"flag": "🇱🇦",
		"calling_code": 856
	},
	{
		"name": "Latvia",
		"flag": "🇱🇻",
		"calling_code": 371
	},
	{
		"name": "Lebanon",
		"flag": "🇱🇧",
		"calling_code": 961
	},
	{
		"name": "Lesotho",
		"flag": "🇱🇸",
		"calling_code": 266
	},
	{
		"name": "Liberia",
		"flag": "🇱🇷",
		"calling_code": 231
	},
	{
		"name": "Libya",
		"flag": "🇱🇾",
		"calling_code": 218
	},
	{
		"name": "Liechtenstein",
		"flag": "🇱🇮",
		"calling_code": 423
	},
	{
		"name": "Lithuania",
		"flag": "🇱🇹",
		"calling_code": 370
	},
	{
		"name": "Luxembourg",
		"flag": "🇱🇺",
		"calling_code": 352
	},
	{
		"name": "Macau",
		"flag": "🇲🇴",
		"calling_code": 853
	},
	{
		"name": "Macedonia",
		"flag": "🇲🇰",
		"calling_code": 389
	},
	{
		"name": "Madagascar",
		"flag": "🇲🇬",
		"calling_code": 261
	},
	{
		"name": "Malawi",
		"flag": "🇲🇼",
		"calling_code": 265
	},
	{
		"name": "Malaysia",
		"flag": "🇲🇾",
		"calling_code": 60
	},
	{
		"name": "Maldives",
		"flag": "🇲🇻",
		"calling_code": 960
	},
	{
		"name": "Mali",
		"flag": "🇲🇱",
		"calling_code": 223
	},
	{
		"name": "Malta",
		"flag": "🇲🇹",
		"calling_code": 356
	},
	{
		"name": "Marshall Islands",
		"flag": "🇲🇭",
		"calling_code": 692
	},
	{
		"name": "Martinique",
		"flag": "🇲🇶",
		"calling_code": 596
	},
	{
		"name": "Mauritania",
		"flag": "🇲🇷",
		"calling_code": 222
	},
	{
		"name": "Mauritius",
		"flag": "🇲🇺",
		"calling_code": 230
	},
	{
		"name": "Mayotte",
		"flag": "🇾🇹",
		"calling_code": 262
	},
	{
		"name": "Mexico",
		"flag": "🇲🇽",
		"calling_code": 52
	},
	{
		"name": "Micronesia",
		"flag": "🇫🇲",
		"calling_code": 691
	},
	{
		"name": "Moldova",
		"flag": "🇲🇩",
		"calling_code": 373
	},
	{
		"name": "Monaco",
		"flag": "🇲🇨",
		"calling_code": 377
	},
	{
		"name": "Mongolia",
		"flag": "🇲🇳",
		"calling_code": 976
	},
	{
		"name": "Montenegro",
		"flag": "🇲🇪",
		"calling_code": 382
	},
	{
		"name": "Montserrat",
		"flag": "🇲🇸",
		"calling_code": 1664
	},
	{
		"name": "Morocco",
		"flag": "🇲🇦",
		"calling_code": 212
	},
	{
		"name": "Mozambique",
		"flag": "🇲🇿",
		"calling_code": 258
	},
	{
		"name": "Myanmar",
		"flag": "🇲🇲",
		"calling_code": 95
	},
	{
		"name": "Namibia",
		"flag": "🇳🇦",
		"calling_code": 264
	},
	{
		"name": "Nauru",
		"flag": "🇳🇷",
		"calling_code": 674
	},
	{
		"name": "Nepal",
		"flag": "🇳🇵",
		"calling_code": 977
	},
	{
		"name": "Netherlands",
		"flag": "🇳🇱",
		"calling_code": 31
	},
	{
		"name": "New Caledonia",
		"flag": "🇳🇨",
		"calling_code": 687
	},
	{
		"name": "New Zealand",
		"flag": "🇳🇿",
		"calling_code": 64
	},
	{
		"name": "Nicaragua",
		"flag": "🇳🇮",
		"calling_code": 505
	},
	{
		"name": "Niger",
		"flag": "🇳🇪",
		"calling_code": 227
	},
	{
		"name": "Nigeria",
		"flag": "🇳🇬",
		"calling_code": 234
	},
	{
		"name": "Niue",
		"flag": "🇳🇺",
		"calling_code": 683
	},
	{
		"name": "Norfolk Island",
		"flag": "🇳🇫",
		"calling_code": 672
	},
	{
		"name": "North Korea",
		"flag": "🇰🇵",
		"calling_code": 850
	},
	{
		"name": "Northern Mariana Islands",
		"flag": "🇲🇵",
		"calling_code": 1670
	},
	{
		"name": "Norway",
		"flag": "🇳🇴",
		"calling_code": 47
	},
	{
		"name": "Oman",
		"flag": "🇴🇲",
		"calling_code": 968
	},
	{
		"name": "Pakistan",
		"flag": "🇵🇰",
		"calling_code": 92
	},
	{
		"name": "Palau",
		"flag": "🇵🇼",
		"calling_code": 680
	},
	{
		"name": "Palestine",
		"flag": "🇵🇸",
		"calling_code": 970
	},
	{
		"name": "Panama",
		"flag": "🇵🇦",
		"calling_code": 507
	},
	{
		"name": "Papua New Guinea",
		"flag": "🇵🇬",
		"calling_code": 675
	},
	{
		"name": "Paraguay",
		"flag": "🇵🇾",
		"calling_code": 595
	},
	{
		"name": "Peru",
		"flag": "🇵🇪",
		"calling_code": 51
	},
	{
		"name": "Philippines",
		"flag": "🇵🇭",
		"calling_code": 63
	},
	{
		"name": "Pitcairn Islands",
		"flag": "🇵🇳",
		"calling_code": 64
	},
	{
		"name": "Poland",
		"flag": "🇵🇱",
		"calling_code": 48
	},
	{
		"name": "Portugal",
		"flag": "🇵🇹",
		"calling_code": 351
	},
	{
		"name": "Qatar",
		"flag": "🇶🇦",
		"calling_code": 974
	},
	{
		"name": "Romania",
		"flag": "🇷🇴",
		"calling_code": 40
	},
	{
		"name": "Russia",
		"flag": "🇷🇺",
		"calling_code": 7
	},
	{
		"name": "Rwanda",
		"flag": "🇷🇼",
		"calling_code": 250
	},
	{
		"name": "Samoa",
		"flag": "🇼🇸",
		"calling_code": 685
	},
	{
		"name": "San Marino",
		"flag": "🇸🇲",
		"calling_code": 378
	},
	{
		"name": "Saudi Arabia",
		"flag": "🇸🇦",
		"calling_code": 966
	},
	{
		"name": "Senegal",
		"flag": "🇸🇳",
		"calling_code": 221
	},
	{
		"name": "Serbia",
		"flag": "🇷🇸",
		"calling_code": 381
	},
	{
		"name": "Seychelles",
		"flag": "🇸🇨",
		"calling_code": 248
	},
	{
		"name": "Sierra Leone",
		"flag": "🇸🇱",
		"calling_code": 232
	},
	{
		"name": "Singapore",
		"flag": "🇸🇬",
		"calling_code": 65
	},
	{
		"name": "Sint Maarten",
		"flag": "🇸🇽",
		"calling_code": 1721
	},
	{
		"name": "Slovakia",
		"flag": "🇸🇰",
		"calling_code": 421
	},
	{
		"name": "Slovenia",
		"flag": "🇸🇮",
		"calling_code": 386
	},
	{
		"name": "Solomon Islands",
		"flag": "🇸🇧",
		"calling_code": 677
	},
	{
		"name": "Somalia",
		"flag": "🇸🇴",
		"calling_code": 252
	},
	{
		"name": "South Africa",
		"flag": "🇿🇦",
		"calling_code": 27
	},
	{
		"name": "South Georgia",
		"flag": "🇬🇸",
		"calling_code": 500
	},
	{
		"name": "South Korea",
		"flag": "🇰🇷",
		"calling_code": 82
	},
	{
		"name": "South Sudan",
		"flag": "🇸🇸",
		"calling_code": 211
	},
	{
		"name": "Spain",
		"flag": "🇪🇸",
		"calling_code": 34
	},
	{
		"name": "Sri Lanka",
		"flag": "🇱🇰",
		"calling_code": 94
	},
	{
		"name": "Sudan",
		"flag": "🇸🇩",
		"calling_code": 249
	},
	{
		"name": "Suriname",
		"flag": "🇸🇷",
		"calling_code": 597
	},
	{
		"name": "Swaziland",
		"flag": "🇸🇿",
		"calling_code": 268
	},
	{
		"name": "Sweden",
		"flag": "🇸🇪",
		"calling_code": 46
	},
	{
		"name": "Switzerland",
		"flag": "🇨🇭",
		"calling_code": 41
	},
	{
		"name": "Syria",
		"flag": "🇸🇾",
		"calling_code": 963
	},
	{
		"name": "Taiwan",
		"flag": "🇹🇼",
		"calling_code": 886
	},
	{
		"name": "Tajikistan",
		"flag": "🇹🇯",
		"calling_code": 992
	},
	{
		"name": "Tanzania",
		"flag": "🇹🇿",
		"calling_code": 255
	},
	{
		"name": "Thailand",
		"flag": "🇹🇭",
		"calling_code": 66
	},
	{
		"name": "Timor-Leste",
		"flag": "🇹🇱",
		"calling_code": 670
	},
	{
		"name": "Togo",
		"flag": "🇹🇬",
		"calling_code": 228
	},
	{
		"name": "Tokelau",
		"flag": "🇹🇰",
		"calling_code": 690
	},
	{
		"name": "Tonga",
		"flag": "🇹🇴",
		"calling_code": 676
	},
	{
		"name": "Tunisia",
		"flag": "🇹🇳",
		"calling_code": 216
	},
	{
		"name": "Turkey",
		"flag": "🇹🇷",
		"calling_code": 90
	},
	{
		"name": "Turkmenistan",
		"flag": "🇹🇲",
		"calling_code": 993
	},
	{
		"name": "Tuvalu",
		"flag": "🇹🇻",
		"calling_code": 688
	},
	{
		"name": "Uganda",
		"flag": "🇺🇬",
		"calling_code": 256
	},
	{
		"name": "Ukraine",
		"flag": "🇺🇦",
		"calling_code": 380
	},
	{
		"name": "United Arab Emirates",
		"flag": "🇦🇪",
		"calling_code": 971
	},
	{
		"name": "United Kingdom",
		"flag": "🇬🇧",
		"calling_code": 44
	},
	{
		"name": "United States",
		"flag": "🇺🇸",
		"calling_code": 1
	},
	{
		"name": "Uruguay",
		"flag": "🇺🇾",
		"calling_code": 598
	},
	{
		"name": "Uzbekistan",
		"flag": "🇺🇿",
		"calling_code": 998
	},
	{
		"name": "Vanuatu",
		"flag": "🇻🇺",
		"calling_code": 678
	},
	{
		"name": "Venezuela",
		"flag": "🇻🇪",
		"calling_code": 58
	},
	{
		"name": "Vietnam",
		"flag": "🇻🇳",
		"calling_code": 84
	},
	{
		"name": "Western Sahara",
		"flag": "🇪🇭",
		"calling_code": 212
	},
	{
		"name": "Yemen",
		"flag": "🇾🇪",
		"calling_code": 967
	},
	{
		"name": "Zambia",
		"flag": "🇿🇲",
		"calling_code": 260
	},
	{
		"name": "Zimbabwe",
		"flag": "🇿🇼",
		"calling_code": 263
	}
]
`
