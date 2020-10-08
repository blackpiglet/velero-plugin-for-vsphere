/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by crds_generate.go; DO NOT EDIT.

package crds

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	apiextinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

var rawCRDs = [][]byte{
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4V\xcdn+7\x0f\xdd\xcfS\x10\xf7[d\x13\x8fo\xf0uQ\xcc\xee\xd6i\x81\xa0\u007fA\x12dStAK\xb4\xcdF#\xa9\"ǩ\xfb\xf4\x85\xa4\x19{\xe2\xf8\x16\xd94\x9b`H\x8a\xa4\xce9\xa4\xdc,\x16\x8b\x06#?S\x12\x0e\xbe\x03\x8cL\u007f)\xf9\xfc%\xed˷\xd2rX\xeeo\x9a\x17\xf6\xb6\x83\xd5 \x1a\xfa\a\x920$C\xb7\xb4a\xcf\xca\xc17=)ZT\xec\x1a\x00\xf4>(f\xb3\xe4O\x00\x13\xbc\xa6\xe0\x1c\xa5Ŗ|\xfb2\xaci=\xb0\xb3\x94J\xf2\xa9\xf4\xfes\xfbM\xfb\xb9\x010\x89\xca\xf1'\xeeI\x14\xfb\u0601\x1f\x9ck\x00<\xf6\xd4\xc1\x1a\xcd\xcb\x10\x13\xc5 \xac!1I[M6\xf1\x9eRk\xbc\xd8\xd8\xee\xfbWLԚ\xd07\x12\xc9\xe4V\xb6)\fq:\xff\xb5\xe0Zel\xbd^\xfb\xbbr\xe0a*x(.Ǣ?^t\xffĢ%$\xba!\xa1\xbb\xd4pq\v\xfb\xed\xe00\xbd\v\xc8\x05ĄH\x1d\xac\xdc J\xa9\x01\x18q*\x8d-F$\xf67\xe8\xe2\x0eoj:\xb3\xa3\x1ek\xdf\x00!\x92\xffr\u007f\xf7\xfc\xff\xc77f\x00Kb\x12G-\x98_\xbdk\x1eX\x00\xc1Բ\x8b҅\x8542\xde\x02\xdci\x8e8Rja}\x00\xddј\an\v\xaa\x80>\x1f\xdaP\"oj\f<z\x8c\xb2\vz\r+\x17<\xfd\x90B?\x99J\xf8-9\xd2\\\xe1\xe9\x98\xed\x81\xe4\xd8\xd6\xd4B)\x8d\xec\xa5T5\x89,yet\xb0\t\tpD\x12NP\x8e\tO\x17\x1c;\xb4Y\xbcT\xb3T)\x80\xeePᕝ\x835\xc1 dA\x03(\xba\x97\xf2\u007fG\xb3\xac\x00\xbfzw8݉Դ\xb0z\x10ؤ\xd0W\x01E4%=*`\xa2\"\x17\xb2\xc0\x1e\xbe8\x17^\xc9\xfer\n\x9aj\xa2\xc9!\xc1_\x03oJ\xc1c\xa2\x8c9\xf8\xa0\x97\xcf\xe7\xd0\x10)\x95\xa9\xa9\xd96Ȯ\xbd:\x92\x1eS\xf6+O\xba\xae\u007fx\x9ei\xee\x04`\xa5\xfe\xcc\x04\xa0\x87,K\xd1\xc4~ۼw`Jx\x98\x978\xed\x967\xd1oU\x98\x85Z\xa3\xde\xf02J\x9e\xec\xa8m\b\x19\x16\x96LD\"!_\xd7L6\xa3\x87\xb0\xfe\x83\x8c\xb6\xf0H)\x1f\x04م\xc1٬\x97=%\x85D&l=\xff}\xcc&\x13\xad\x0e\x95$C\xab\x94<:أ\x1b躈\xb2Ǭ\xbc\x9c\x17\x06?\xcbPB\xa4\x85\x9fC\"`\xbf\t\x1d\xecT\xa3t\xcb\xe5\x96uڛ&\xf4\xfd\xe0Y\x0f\xcb2/\xbc\x1e4$YZړ[\no\x17\x98̎\x95\x8c\x0e\x89\x96\x18yQ\x9a\xf5ew\xb6\xbd\xfd\xdf$z\xb9\xba\x00\xf5;\x0e\xd6g\xa3\xbcr\xc8}\xf7\x91\x93e\xcd\xfd\v?y\xcf\xd5\xc5P\x8f\xd6\xfb\x9fhȦ\x8c\xe4\xc3\xf7\x8fO\xa7I-TUVN\xa1r\"(\x83\xcb~S&\x8f\xc7\xd1\xc9Y\xc8\xdb\x18\xd8k\x9dp\xc7\xe4\x15dX\xf7\xac\x99\xf9?\a\x12\xcdܵ\xb0*OM\x99\xd6hQɶp\xe7a\x85=\xb9\x15\n\xfd\xe7\xf4d4e\x91\xc1\xfb\x18A\xf3W\xf2<\xb8\xe24s\xe4e\x13+\x91\xf7\x98\xb0'\xa5t6\x8dhmy~\xd1\xdd_\x9c\xef\x0f\xcc\xebŲ\xf3E\xf9!\xfd\xc8ޜ\xbf#y\xa9|\xe0l\xe6\x93\x13\xcdԷx\xbf\x96f\xbe\x8b\x1a\x9f\xf9/\xa2v\xe6\x9f_\xaf\xf9*\x1c\x92Uj;\xd04P5hH\xb8\xa5\xd1\"\x8a:\x14\xb4\xd1\x18\x8a:\xf6;\xff\xd9\xf0\xe9ӛ_\x01\xe5\xd3\x04_9\x93\x0e~\xfb\xbd\xa9Y\xc9>O\x8f{6\xfe\x13\x00\x00\xff\xff\xb3\x96\xf8\xea\x95\t\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4U?s#\xb7\x0f\xed\xf7S`\xeeW\\sZ\x9d\xe7\x97\"\xb3\x9d\xa3Kq\x93?\xe39{\xdcdR@$$!\xe6\x92\f\x80\x95\xe2|\xfa\fII\x96e)\xe3&\xdb-\x00\x02\x0fx\x8f`7\x9b\xcd:\xcc\xfcH\xa2\x9c\xe2\x00\x98\x99\xfe2\x8a\xe5O\xfb\xa7\xef\xb5\xe74\xdf\xdetO\x1c\xfd\x00\x8bI-\x8d\xdfH\xd3$\x8e\xbeЊ#\x1b\xa7؍d\xe8\xd1p\xe8\x000\xc6dX\xccZ~\x01\\\x8a&)\x04\x92ٚb\xff4-i9q\xf0$5\xf9\xa1\xf4\xf6s\xff]\xff\xb9\x03pB\xf5\xf8\x03\x8f\xa4\x86c\x1e N!t\x00\x11G\x1a`\x89\xeei\xcaB9)[\x92g\x17\x90G\xed\x9b\xd9\voIz\x17\xd5\xe7~;\xeeP\xa8wi\xec4\x93+p֒\xa6|\xc8q-\xb8U\xda\xc3o\xad\xffP\x0f|;\x16]\x94\xa2\xd5\x1fX\xed\xa7\xeb1?\xb3Z\x8d\xcba\x12\f\xd7\xe0\xd7\x10帞\x02ʕ\xa0\x0e@]\xca4\xc0\xaf\x05^FG\xbe\x03\xd8O\xb0\u009d\xedg\xb4\xbd\xc1\x907x\xd3Һ\r\x8dغ\x01H\x99\xe2\xed\xdd\xd7\xc7\xff߿2\x03xR'\x9c\xad\xb2\xf1\xf1r7\xc0\n\x93\x92\aK\xe0\v\xff4G\xe7H\x15\xf0́\x1e\xe0\x16\"\xed\xde8`\xc7!\xc0\x92\x1a\xd3\xe4\x01vl\x1b\xb0\r\xc1KЗ\xca\xcd'X\by\x8a\xc6\x18\x00\xa3\x87\xdb\x10Ҏ\xfcq\x00ڒ\x11ۆ\xa4\xe4,Y\xe2\xc1\v\xb6A\xab\xa6s\f\xf7\x99\x1c\xc0\x0e\xf5\b\x82#$\xa9\xb1ok\x14\xf5\xf0\x8a[Եt=\xc0\xc3\x05\x17\xac\x98\x82o0\v\xc0)\xfbZ\xef\xd8sA\viu1\xef\x01]\xff\xf1HS\x96\x94I\x8c\x0f\xfal\x1f\x9ec>u\x02\xb0\xd1xf\x02\xb0\xe7\xa2%5\xe1\xb8\xee\xde:P\x04\x9fOK\xbc\xec\x89WѯuS\xa4բ\xf6\x02\xd1\xda\xd8^\xa4\xe4\xf7jl\r\xb3\x82P\x16R\x8ame\x143FH\xcb?\xc8Y\x0f\xf7$\xe5 \xe8&M\xc1\x97M\xb2%1\x10ri\x1d\xf9\xefc6-z,e\x02\x1a\xa9\x01G#\x89\x18`\x8ba\xa2OU9#>\x83P\xc9\vS<\xc9PC\xb4\x87_\x92\x10p\\\xa5\x016fY\x87\xf9|\xcdv\u0601.\x8d\xe3\x14ٞ\xe7u\x9d\xf1r\xb2$:\xf7\xb4\xa50W^\xcfP܆\x8d\x9cMBs\xcc<\xab`c݃\xfd\xe8\xff'\xfb\xad\xa9\x1f/\x8c\xfa\r\a\xcb3\x1d\f\xef9T7տPS\xb6T\xb9\xc0\xb8?\xdaZ\u007fa\xa0\x98\xea%\xfc\xf1\xfe\x01\x0ex\x1bK\x8d\x90\x97P}\xe1\xa6̕㊤E\xae$\x8d5\vE\x9f\x13\xc7v\xfd\\`\x8a\x06:-G\xb6B\xfa\x9f\x13\xa9\x15\xdazX\xd4\x17\xe3\xe4j\xf4\xf05\xc2\x02G\n\vT\xfaϙ)\xd3\xd4Y\x19\xde\xfb\xb89}\xec\u0383ۜN\x1ce\x85\xe7\xc6\xe1\x1d\n\x8ed$g\x17\x11\xbd\xaf\xaf(\x86\xbb\x8bW\xfb\x1dW\xf5b\xd9\xd3=\xfa\x0e\xfd\x14NX\xe8DA\xb3\xcb\xe8\xcf\xfc\xa7e\xba\xab\xb0\xb4\xa8\xc5\x0f`2Q3X\x12\\\xd3ޢ\x866ծ\xcbs\x92m\xbf\xc9N_\xe0\x0f\x1f^=\xa3\xf5ץ\xd8f\xa7\x03\xfc\xf6{ײ\x92\u007f<\xbc\x88\xc5\xf8O\x00\x00\x00\xff\xff(|\x8f\x15\xe4\b\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4WM\x8f\xdb6\x13\xbe\xfbW\f\xf2\x1e\xf6-\x10\xcb\t\xdaC\xe1[\xe0mZ\xa3M\xba\xc8\x06{)z\xa0ȱŮD\xb2\x9c\x91R\xb7\xe8\u007f/\x86\x94\xfc!\xd9\xebl?t\x139\x9f\xcf\xcc<$g\xf3\xf9|\xa6\x82}\xc0Hֻ%\xa8`\xf17F'\u007fT<~M\x85\xf5\x8b\xee\xf5\xec\xd1:\xb3\x84UK\xec\x9b\x0fH\xbe\x8d\x1aoqc\x9de\xebݬAVF\xb1Z\xce\x00\x94s\x9e\x95,\x93\xfc\x02h\xef8\xfa\xba\xc68ߢ+\x1e\xdb\x12\xcb\xd6\xd6\x06c2>\xb8\xee^\x15_\x15\xaff\x00:bR\xffh\x1b$VMX\x82k\xebz\x06\xe0T\x83Kеw\xb8\x89\xbe!\xa7\x02U\x9e\xa9(\x95~l\x83\x89\xb6\xc3XhG&\x14]\xf3IE,\xb4of\x14PK(\xdb\xe8۰\x84\xa7\x85\xb3\x97>\xf4>mq\xf86\xfa\xe6\xbew\x98\xf6jK\xfc\xfd\xf9\xfd\x1f,e\x99P\xb7Q\xd5\xe7BN\xdbdݶ\xadU<#0\x03 \xed\x03.ὄ\x13\x94F3\x03\xe8\xd1J\xe1\xcd{<\xbaת\x0e\x95z\x9dM\xea\n\x1b\x95\xa3\a\xf0\x01ݛ\xbb\xf5×\xf7'\xcb\x00\x06IG\x1b8!\u007f3\xcd\x00,AKh\x80}\xae\a\x82\x02\x87\x9f \xf6Ň\xff\xf3.X\xad\xeaz\a\n\xee\x1eV_\x80\x84\x0f\n\x86\f\n\x80\x1f\x9dF\xe0\na0{sCpW)B\xa8\x14\x014\xbe\xcb.\x86}F\x0369\xefTm\x9f\xf0\x9e|\x89\xe5\xc1\x1b\xacoo\xf6م\xe8\x03F\xb6C\x19\xf3w\xd4\xe7G\xabc,\x04\xae,\x05F\x1a\x1c)\xf9\xe9\x81G\xd3#\f~\x03\\Y\x82\x88!\"\xa1\xcb-/\xcbʁ/\u007fA\xcd\x05\xdcc\x14E\xa0ʷ\xb5\x91I\xe802D\xd4~\xeb\xec\xef{k$\x99\x8a\x9bZ1\x12\x83u\x8cѩZ`h\xf1%(g\xa0Q;\x88(v\xa1uG\x16\x92\b\x15\xf0\xceG\x04\xeb6~\t\x15s\xa0\xe5b\xb1\xb5<̰\xf6M\xd3:˻E\x1aG[\xb6\xec#-\fvX/\xc8n\xe7*\xea\xca2jn#.T\xb0\xf3\x14\xacKs\\4\xe6\u007f\x03\xf4ts\x02\x1e\xef\xa4G\x89\xa3uۣ\x8d48O\xa0,\x83#eV\xbdj\xce\xe2\x00\xa6,\t\x1e\x1f\xbe\xb9\xffx\xa8z\x02<c{\x10\xa5\x03\xcc\x02\x91u\x1b\x8cYr\xdf$\xe8L\xf0\xd6q\xfaѵE\xc7@m\xd9X\x96\xfa\xfd\xda\"\xb1T\xa0\x80U\"/(\x11\xda`\x14\xa3)`\xed`\xa5\x1a\xacW\x8a\xf0?\aYФ\xb9\x80\xf7y0\x1f\xf3\xeeX8\xe3t\xb41\xd0\xe0\x85\x9aL8\xe0>\xa0NJvc\x91\x0em-\xbdZb\xa6,3\x9ezX\xdf\x16\x00\x1f+\x84w}l\xa9*%\x82\xef0Fk\f\xba\x97\xa9\x0e\x1b\x1f\x1b\xc5y\x8ep\x9f\t\x1c*ܻ\xd6\x05\xc0\x9b\xbb\xf5\xb7B\xdei\x10R\xef\xe4\xcd]ҕ|\xc5\xce!\xbcL\x1a\xc5I\xb2\xe7I\xa1'\x86d}\xbc>\x02h\x1fD\x1f\xf2\xbe-K\x94v\xcd>\xcd\xc4ƅ\xd2ɗO\xa2\x0f\x18<Y\xf6qw%\x00A5\xabH\xf7\xf7:\x92nD\x8e\x16;<eD\xa9L_\v7\x9c!'l\xbc\xb8{XAm;$\xb0\x0e\x9a\x96\x18*\xd5!(\xad\x91\xf6\x94tp\xf5\x9c\xd4Rw\xac\x94\xd3X_\xc9j\x88&\v\x83u\xc6ja\xc1a2S=\xf3\x9ew[/P\x1f\x1d1#m\xad\x9c\xf4\x1a!\x83bPnǶ\x91\x02mdnOЉ\xa8t%m\r\x8c\xb1\xb1B\xb6\xa1J3\x0e\xebͩ\xa8\x9cUY\xdcL\xc4/`Rz_\xa3r\xa3\xdd)+N\xd0\x18\x88\xf1\xb8\xaf\xffy\xa3\x9d\xe7\t\xf9\xf2\x14.\xa1\xdc\xf1\xa5\\\xceZ\x1c\xc0Y\xdfNm^T\x93\x8aڈ#\f\xe6\xfb\x01\x1c-\x8f\xc7c\xb4}\xd4b\xa3\x1d\xc1y\xb4t\x88\xf7\xb3\xa8\x92\x15\xb7#\x9e\xb8\xcc\x1f\r\x12\xa9-^\xa9\xec\xbb,\x95O\xbd^\x05T\xe9\xdb\xe1X\xf2\x0eo\xa8w]<\xa7\x16\xa9\x11\xafxϷ\xae\xbe\xb1t\x1bc:\x03Y\xaeV=\x9dM\xf8\xffY1\f=\xfa\x9dr\xa6\xbe\x16\x8cPR\x95\x04'\\ʕb\xf8\xa4\xe8\xe4x9\x9eƉ\xe5\xcbu\x81'\xb9}\x12՞\xdf{\x94\xd2sA\x06\xe4\x1c\xdbG\xdc`D\xa7\xd3\x15as\xa2+7\x88\xe1\xd84\xf9\xac\xdb\xff\xe6\xf9NT[\xcaE\"\x97C\xc8\xe9\xcd\xdd:{,\u0b4fB]\xe0\xb9\xcaW\x99h\xe6AEޥ\x02\xd0\xcb\x13o\xc3LM\xabu\xa5bp\x91\x90&\xc8|\x06)\x1d\xf0\xf8;q\xa4\x87\xcc\xf58\xe4)4\xc4!*\xffr\x1c\xe7\xe9\t\xceSJ^\x96(.\x8cɄT.lL\xbd\xceS\xbf\xcc.j\x91\\w\xcd\x128\xb6\xd99\xb1\x8fB@G+m\xb9\xbf\xb0\x0f\x86{R\x83?\xfe\x9c\x1d\xf8M\xce\xfa\xc0hޏ\x1f\xbd/^\x9c\xbc`ӯ\xf6\xce\xd8\xfc\xb2\x87\x9f~\x9ee\xc7h\x1e\x86G\xa9,\xfe\x15\x00\x00\xff\xffo-bHS\x10\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4V=o\xe46\x10\xed\xf5+\x06\x97\xc2M\xa4=#)\x02u\xc1:\x85\x91\xdc\xc1\xb0\r7A\n.9^1\xa6Hff\xa8\xc4\t\xf2\xdf\x03\x92\x92\xb5\u07b5}\xbe\"\xdbi>\x9fޛ\x19mӶm\xa3\xa2\xbdCb\x1b|\x0f*Z\xfcK\xd0\xe7'\xee\x1e~\xe0Ά\xcdt\xde<Xoz\xd8&\x960^#\x87D\x1a/\xf0\xdez+6\xf8fDQF\x89\xea\x1b\x00\xe5}\x10\x95͜\x1f\x01t\xf0B\xc19\xa4v\x8f\xbe{H;\xdc%\xeb\fR)\xbe\xb4\x9e>v\xdfw\x1f\x1b\x00MX\xd2o\xed\x88,j\x8c=\xf8\xe4\\\x03\xe0Ո=\x18t(\xc8^E\x1e\x82p\xb7S\xfa!ECvB\xea\xb4g\x13\xbbi\xfcS\x11v:\x8c\rG\xd4\x19ǞB\x8a=\xbc\x1d\\[̸\xeb;_\x94n7s\xb7\xe2p\x96\xe5\xe7\x17\x9c\xbfX\xae\x01\xd1%R\xee\x04i\xf1\xb1\xf5\xfb\xe4\x14\x1d{\x1b\x00\xd6!b\x0f\x9f3\x84\xa84\x9a\x06`\xa6\xa7@jg\x02\xa6s\xe5\xe2\xa0\xcek==\xe0\xa8*b\x80\x10\xd1\xffxuy\xf7\xdd\xcd33@\xa4\x10\x91\xc4./W\u007f\a\xd2\x1fX\x01\f\xb2&\x1b\xa5\bs\x96\v\xd6(0Ysd\x90\x01\x17hhf\f\x10\xeeA\x06\xcb@\x18\t\x19}\x9d\x82lV\x1e\xc2\xeew\xd4\xd2\xc1\rRN\x04\x1eBr&\x0fǄ$@\xa8\xc3\xdeۿ\x9f\xaa1H(m\x9c\x12d\x01\xeb\x05\xc9+\a\x93r\t\xbf\x05\xe5\r\x8c\xea\x11\bs]H\xfe\xa0B\t\xe1\x0e>\x05B\xb0\xfe>\xf40\x88D\xee7\x9b\xbd\x95e\xacu\x18\xc7\xe4\xad<nʄ\xda]\x92@\xbc18\xa1۰ݷ\x8a\xf4`\x05\xb5$\u008d\x8a\xb6-`}\x19\xedn4\xdfм\b|\xf6\x8c<y\xcc*\xb2\x90\xf5\xfb\x03G\x19\xa77X\xce\x13\x05\x96Aͩ\xf5-V2\xb3)\xf3q\xfd\xd3\xcd-,\xad+\xe1\x95\xdb5\x94W\x9a3E\xd6\xdf#\xd5\xc8{\nc\xa9\x82\xde\xc4`\xbd\x94\a\xed,z\x01N\xbb\xd1J\xd6\uf3c4,Y\x81\x0e\xb6e\x9fa\x87\x90\xa2Q\x82\xa6\x83K\x0f[5\xa2\xdb*\xc6\xff\x9d\xe4\xcc&\xb7\x99\xbc\xf7\xd1|x\x8a\x8e\x83+O\a\x8e\xe58\xbc\xa2\xc9MD\x9d%)\x1c\x95۷\x12\x9fS\x9fe\xbe\xbca\xf9Wo\xce5\xc6\xc0V\x02=\x1e\xfb\x8f\xba\xde\x0e8\xa7dE眼\r\x84B\x16',x\x96\xbbQ$\xedJ\x92_\x0eG\tX\xae\xd2\xe6\xean\v\xceN\xc8`=\x8c\x89\x05\x065!(\xad\x91\x9f\xb6l\xedt\x02\xee\x15\xa2\v\u007fs\x8fˋ\xd3wz5-\x8f\x97%<Z\x86\xf6\x84\xa6#\xf7\xda\xeb]ʊ\x92\xc4oh\xbbMDe\xe8K`\xbd]8\xdf\xe4\x95\xdc,i\xb9b\xef\x94zDf\xb5\xc7/(\xfc\xa9F\xd5m\x9fS@\xedB\x92\x97@\x9c\xf1\f\xb2\xfb\x1ai\xe2\xa0\xf8K8\xaer\xcc:\xe0+!\xb8\xf0Q?pO\xb3\xf4\x15\b^\xd4\xe5T\xfa\xf6\xf9\"\x9ddq\xbec\xa6\a\xa1\x84\xd5 \x812\xc3\a\x96\xb4{\xba\xc4K\xe1Y~\xf8\xe7\xdff\x9d\x84<\xf1Q\xd0|>\xfe\xc6\u007f\xf8\xf0\xec\xb3]\x1eu\xf0\xc6\xd6\u007f1\xf0\xeboMm\x8c\xe6n\xf9\x1eg\xe3\u007f\x01\x00\x00\xff\xff\x827!\xdd?\t\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4X\xcdo\x1b\xb7\x12\xbf\xeb\xaf\x18\xf8\x1d|\xb1V\t^\xf0\xf0\xa0[j\xf7\xc3H\x13\x18\xb6\xe1K\xd1\x03\x97\x1ciY\xef\x92,g\xa8D-\xfa\xbf\x17C\xeeJZ}\xd9)\x1aݖ;\x9f\xbf\x99\xf9\rW\x93\xe9t:Q\xc1>a$\xeb\xdd\x1cT\xb0\xf8\x85\xd1\xc9\x13U\xcf\xff\xa7\xca\xfa\xd9\xea\xed\xe4\xd9:3\x87\xebD\xec\xbb{$\x9f\xa2\xc6\x1b\\Xg\xd9z7鐕Q\xac\xe6\x13\x00\xe5\x9cg%\xc7$\x8f\x00\xda;\x8e\xbem1N\x97\xe8\xaa\xe7Tc\x9dlk0f\xe3\x83\xeb՛\xea]\xf5f\x02\xa0#f\xf5G\xdb!\xb1\xea\xc2\x1c\\j\xdb\t\x80S\x1d\u0381\x9c\n\xd4x\xa6\xaaV\xfa9\x05\x13\xed\nc\xa5\x1d\x99P\xad\xba\xcf*b\xa5}7\xa1\x80Z\"XF\x9f\xc2\x1c\xce\v\x17\xe3}\xc4%ۇ\xdeO>j-\xf1\x87\xd1\xf1ϖʫЦ\xa8ڝ\xb8\xf2)Y\xb7L\xad\x8a\xdb\xf3\t\x00i\x1fp\x0e\x9f\xc4UP\x1a\xcd\x04\xa0\a \xbb\x9e\xf6)\xaeު64\xeam\xb1\xa4\x1b\xecT\x89\f\xc0\at\xef\xefn\x9f\xfe\xfb0:\x060H:\xda\xc0\x19\xcc\xcbM\x98`\t\x12\xa1\x01\xf6\x10\xf1\xf7\x84\xc4\xc0\x8dbP\x9b\xc0D\x84\xd53\xba\n\xe06?9\xcf\x1b\xa5N9\xb5D\xe0\x06\xc1\xba\x15:\xf6q\r~\xb1M\x17\x943`<\x165pX\xf4\xf0\x8b%\x06\xeb\xc0G\x83QNt\xeb]1\x14\xfb\x16\x82E\xf4\xddN$\x97\x9blB\xf4\x01#ۡ$\xe5\xb7Ӫ;\xa7\xfb\xb9\v<E\n\x8c\xf4(Rv\xda\x03\x8d\xa6GT\x92\xe0\xc6\x12D\f\x11\t]\xe9Z9V\x0e|\xfd\x1bj\xae\xe0\x01\xa3(\x025>\xb5F\x9ay\x85\x91!\xa2\xf6Kg\xff\xd8X#\xc9Pܴ\x8a1'\xce\x18\x9dja\xa5ڄW\x19\xa4N\xad!\xa2\u0605\xe4v,d\x11\xaa\u08cf\x02\xf2\xc2ϡa\x0e4\x9f͖\x96\x871Ծ뒳\xbc\x9e割ub\x1fifp\x85\xed\x8c\xecr\xaa\xa2n,\xa3\xe6\x14q\xa6\x82\x9d\xe6`]\x1eŪ3\xff\x19P\xa7\xcb\x11x\xbc\x96\x9e$\x8e\xd6-w^\xe4!8\x83\xb2L\x83\xf4\x8a\xeaUK\x16[0\xe5H\xf0\xb8\xff\xfe\xe1q[\xf0\fx\xc1v+J[\x98\x05\"\xeb\x16\xd20\"\x99\xfbC\xac\xa03\xc1[\xc7\xf9A\xb7\x16\x1d\x03\xa5\xba\xb3LC[K\x05*\xb8\xce\xfc\x035B\nF1\x9a\nn\x1d\\\xab\x0e\xdbkE\xf8\xcdA\x164i*\xe0\xbd\x0e\xe6]\xea\xdc\x17.8\xed\xbc\x18(\xedDM\x1e\x02\xea<̂Q\xe6\xea-\xf0\xa2:\xd2<>a\xf2+Ly\x8f\xc1\x93\x95i\xdf\u007f\xbf\xe7\xf5\xb1\xc1^E*\xda\xeb\xc84l\xf9\xc5Ie\xb2\xa0\x1b\xa8/\a9\x90\xd4\xec\xee\xe9\x1aZ\xbbB\x12\xc2\xe8\x1214j\x85\xa0\xb4F\xdaL\xd6\xd6\xfaA@'\xc0\x95߀\xc0Oʙ\x16_\xc8\xe5~$\f\x11\x17Ҕ\xecA\xc1\x87Tct\xc8H\x1b\x93W\xa0S\x8c\xe8\xb8]\x83\x02ɡN\\\x1aW\xfa\xb8F\xc8\xcbԠ\x91\xb4$\x85E\x92\xbe9\x88\xe1t-\xa0\xb0ޏy\x87\x1dy\xb7\x17\xff\xfb\xbb\xdb,:tA\xde}\xb0\xf0qL\xbb5\xcat\xe6\xec\xd0\xe9<#\x8b\x91\xae\x8c\x90t\x8c]X4WYy\xf3\by\xf2s\x91j\x1c\x12\xd32U\xef\xefn\x8b\xc7\n~\xf0\x11\x94[\x83\xe7\xa6\xccr4Ӡ\"\xafs\xa9\xe8j\xe4M\x06\xd8F4\xd5\xd1\x04\xcf\xd4\x16\x8e\xf2\xd4Qd\x06\xba\x92`Ţ\xd0\xfcI<\xfeI\x1cys\xbf\x1c\x87\xec\xfe!\x0eQ\xf9\x97\xe3\x18\xa0<\x8cd\x9a\x91:r,Q\x9c\x18\xa8\x03\x02\x92\xdf0\xd6\xd7\xcail_\x18\xa8\x87\x910Xg\xacVe\x86\xfa\xbb\x88\a]\xdey\xb7\xf4\x92\xfe`\xbf\x82}m\xad\x9c\xb4\x1c!\x83\\`ܚm'\xa0-\xa4\xf9r\x8f\x0e\x8c\x13Q\xe9\x06e91\xc6\xce\xca\x1e\x0eM\xa6\u007f\xe9\xf4\x91h\xa3\xa8\x177\a\xe2'`\xa9\xbdoQ\xb9\xc9\xcb\xc0O\x0f\xc8t\xef\xf5\x98\x9d^\xb5\x03Xq\xa23[\xe0\xbapR/Xn9;\xf9\n\xd3\xe4\x8b\xce+\xb7A\x87Dj\xf9\x12q~,R\xe5BЫ\x80\xaa}\xe2\x91\xf7K\xea\xc3:\x1c\xf53m}|G\x1e\x89\xa1\x88m\xf8n\xf0\xcah\x8e\xb72\x88h\xa7x\x0e\xf5\x9aOU\xfbhH\xb9=^\x88\xe7Nd\xb6\xfbx[\x15\x1c\x8a2\xf4\xf7W\xa1\x11\xa2_F\xa4#\xabb\xec\xbd\x17۠\x91B\xeb\xd5!\x01\x9c\xdf=\x82\v\xddxw\x82\xda\x06\xf8\xac\xe3\xff\xbd;\xc3Vr\x1b^b<\"\xc1\x9eU\xfb\x9dx\xf96\x1e^Ad\xb77\xaf$1\xb8\xbd)\xdfN\xc2\x195\xa2\xdb|6=\xca\xce\xffl\xdbV\xf8ia۶\xac\xfc\xcf\r\x96\xf5\x98\xdb\x05\x96\xf2\x91\xc4\x1e.\x1e\xb6\x8dy\xf15\xa5\xa7\x95\x1eT?\x1d]7\xa3\xb0e\a/3\xc9\xea6\x11c\xa4+ \xb9V\xed\xee\x9e~yG\xa4\xe0\x9d\xc9\xfc\x9b\x02ƕ%\x1f\a\xbd-y\x88VU\xb2\xdd\xf9\xae\xe4\xa8\xf4\xf3N\x8f\rԳ\xb9\xbd\x1f\x9a\xfc\x8a\x8e?Z\xc1C\xba\x9d\x8e\xaf\xb8\aZ$_\x18f\x0e\x1cS\x99tb\x1f\x85\xd8vNR\xbd\xf9F\x1a\f\xf7t\v\u007f\xfe5\xd92\xaf\xdcK\x03\xa3\xf9\xb4\xff\x9f\xc1\xc5\xc5\xe8/\x81\xfc\xa8\x05\xd7\xf2\u007f\b\xfc\xf2\xeb\xa48F\xf34|\xf7\xcb\xe1\xdf\x01\x00\x00\xff\xff\xc7\x11Gq\x89\x11\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcX\xcdo\xe3\xba\x11\xbf\xfb\xaf\x18\xbc\x1er\xa9\xe5\xb7}\x8b\xa2\xf0\xad\x9b\xec\x02A\xbbA\x10os)z\xa0ȱņ\"\xb5\x9c\xa1\xbd\xee__\f)\xd9\xf2G\xe2\xb8\xe8>_\x82P\xf3\xc5\xdf\xfc\xe6C\x9aL\xa7Ӊ\xea\xec3F\xb2\xc1\xcfAu\x16\u007f0z\xf9\x8f\xaa\x97\xbfPe\xc3l\xfda\xf2b\xbd\x99\xc3m\"\x0e\xed\x13RHQ\xe3\x1d.\xad\xb7l\x83\x9f\xb4\xc8\xca(V\xf3\t\x80\xf2>\xb0\x92c\x92\u007f\x01t\xf0\x1c\x83s\x18\xa7+\xf4\xd5K\xaa\xb1N\xd6\x19\x8c\xd9\xf8\xe0z\xfdk\xf5\xb1\xfau\x02\xa0#f\xf5o\xb6Eb\xd5vs\xf0ɹ\t\x80W-\xce\xc1\x84\x8dwA\x19\xaa\xc4e\x1b\xd6\x18+\xed\xc9tպݨ\x88\x95\x0e\xed\x84:\xd4\xe2~\x15C\xea\xe6\xf0\x86d1\xdb\xc7Z\xeey\xd7{\xc8G\xce\x12\xff\xed\xe0\xf8\xef\x968?\xea\\\x8aʍ\"ʧd\xfd*9\x15\xf7\xe7\x13\x00ҡ\xc39<\x88\xabNi\x94\xb3\xfe\xea\xd9\xf5\xb4\xbf\xdc\xfa\x83r]\xa3>\x14K\xba\xc1V\x95\xc8\x00B\x87\xfe\xaf\x8f\xf7Ͽ-\x0e\x8e\x01\f\x92\x8e\xb6\xe3\f\xe3\x10d\u007fZ#(X\xa3\xc3\x18\xa6\x9dK+\xeb!\"q\x88\xb8S\xefb\xe80\xb2\x1d0(\xbf\x11+F\xa7G\xcen$\x9e\"\x05F\xe8\x80\x04\xdc\xe0p34\xfd\x15 ,\x81\x1bK\x10\xb1\x8bH\xe8\vA\xe4Xy\b\xf5\xbfQs\x05\v\x8c\xa2\bԄ\xe4\x8c\xf0f\x8d\x91!\xa2\x0e+o\xff\xb3\xb3F\xc0!\xbbq\x8a\x91\x18\xacg\x8c^9X+\x97\xf0\x8f\xa0\xbc\x81Vm!\xa2\u0605\xe4G\x16\xb2\bU\xf05D\x04\xeb\x97a\x0e\rsG\xf3\xd9ley`\xbc\x0em\x9b\xbc\xe5\xed,\x93\xd7։C\xa4\x99\xc15\xba\x19\xd9\xd5TE\xddXF\xcd)\xe2Luv\x9a\x83\xf5\x99\xf5Uk\xfe\x10\xfb\x1a\xa1\x9b\x03\xf0x+$ \x8e֯F\x0f2\xeb\xde@Y\xe8\a\x96@\xf5\xaa\xe5\x16{0\xe5H\xf0x\xfa\xbc\xf8\x06\x83\xeb\x02x\xc1v/J{\x98\x05\"\xeb\x97\x18\x8b\xe42\x866[Ao\xba`=\xe7\u007f\xb4\xb3\xe8\x19(խe\xc9\xdf\xf7\x84Ē\x81\nns\xa9C\x8d\x90:\xa3\x18M\x05\xf7\x1enU\x8b\xeeV\x11\xfet\x90\x05M\x9a\nx\xef\x83yܥ\x8e\x85\vN\xa3\aC\x03y%'\x8b\x0e\xb5\xa4$c\x94\xdb\xe2\x1exQ=\xd0<_a\xf2\xab\x95~I\xdd\x13v\x81,\x87\xb8\x95\xfep,s\xe4\xf9ӑ\x8aX_[\x83\xd4\x1b\x93\\\x0f\x8f\x04{X\x86\xb8\xebDՉ\xbax\x1c.\"-\xa8\x94*\x9e\xc8U'Q\xbd\x82\xb2\xfc\xb4\v\x1e\x85P\v\xaf:j\x02?\xe1\x12#z}\xe9r\xb7\xa2\xf8\xe5\x9c\xe28\xc6\xdcAs\x99\xef\x1cQ/_\xee WΤ\x1e\xee=\x10\xb7\x82oM~\xdc*\x16\x8b'\xfev\rzv\xf6\x11\xdcg\xb5Dh\xa4\t\x15\xde\xe7\xb8v\x9e\x88\x15'\x02\xeb\xfb\xfa9\n\xf0*\x1c\xbb\x18\x84\xffh>{\xb6\xbc\xbd\xbf\xbb\x00\xdf\xe3\xb1\xfc\x80\x9a5R7K\x8b\xb1\xc7\x06\xf7\xb6\x01\xb3\xb0@cI\x14<\xa2)\xf7\x93\x89\xb9\x89\x96\x05l\xc0\x1f\x96r\xabY\a\x97Z\xbc\xea\"\xfd\xccُ\xf4\xb7\xef\xf1t$\x9eg@4\xe5.l\xdb\x02yo\x146\x8a@+\xe7\xa4\x01Iz)7\xb8\x1b*\x92C\xba\xe4\xdeC*w\x86O\xc2(\xd4\xc8\v\x03NE\xff\x9a[\x0e9\xbe\x98\xa7\xc5N\xf0\x8d\x04\xed(\xddW\xe4հ\xbf\xd6\xd92A\xdf\xe8m\xc3\n\xb1\xe8\x99<t\xb9\x18\xf3((\xa72\xbaw\x92\xd5;\u06dd\x0em\xe7\xf0p\xb9\xbb\xd0\x10N5Nɠ\xfc\xbe\xfc2\x19\x8a\x92\xf0a\xaf\xbfcCQ\x17ޯ\xd1C\xf0\xb0T֡\x19\xed\x95\x17Xt&&\xfa\x1f\x88$K\xad\xaa\x1d\u0381c\xba\x8ag-\x12\xa9ե^\xfa\xb5H\x95šW\x01U\x87\xc4\a\r\xeb\x86\xfa\x94^U\xd0\x1e\u007f\xf0\x13rܾ7\x91\x0f'\nÞW\xe3.\x93\xe5\x9c\x1bi\xce\xdeX=4W\xf1\x06Q\xb4s\x02\x8e\x12\x06\xb7O\x15\xfc\xa3o\xc9K\xeb\x18#\x1c\xdfr\x18\x00\xb0i\xacn\x84!\x98{t\x8dKi!#\a\x12\xc7)\x14?3\x9d]\xa3\xe8R2\x1fE\xe6\\)\xee\x06\xf6\xf9Z\x94\x1f\xfaԞ\x9a\x9f\xc2\x03nΜ\xde\xfb\xc7\x18V\x11\xe9\x94\xd3Ӂ\xfah\xce<\xcb\xd9=s\xfe%g\xebʱ\xa7\x91\xe4=\xea!\x98K\xc8H\xad\xde)V_\x95W+\x8c\xe0\x83\xc1B\xa1F\x11tV\xbf\xa0\x81\xd4\x1d`\x94Y\xb4\xf7ҏ\xbe\x8dun\xb4͂\"\xa0\x10\xbc\xfc=P\xb6c\xb3ǖ\xee\xfbl\x8c\"\xd2Bs\u007fÃ\xdca\x18\x14Z\x19e\x8a\x82\a˻ \xf6\x1e\xea-(\x1f\xb8\xe9\xefv\xed\x02\x91Syyo\xc8b\xd0\x0474\xd6\xc0ʁOm-մ\x84z\xcbH\x87\xa3(o`c>\ue947ќ\xc3g\xa4\x1ea\xad|\x06\xb8\xafUc\xa9sj\xbb\x8b2\xef\xabRgҮ\xf7\x9dj0&S%?;\x85\xe0\xf5\x91#\xbf\x1c\xc2]\xf0g\xa8\x04\xa3ڶ\x9e\xff\xfc\xf1\xacDAX\xde3W\x18\xcfHd\xb0>\x89\x97\x9f\xe3\xe1\x95I\x0ey\xb3⸽\r\xc9\xf3ŝj\x10<\x18\xa0\xe3\x9c\r\xed\x8f21#N\x951\xb2\xf0\xa9\xa1\xe3\xde\xed;n\xdfH\x87V\x9c\x841\xe0\x917!\xbe\x80%J\x98\xdfG\xe4\xf4{\u0084}\x87\x16É\xe4\x9d3*\xfd\x92\x8d{\x03\x06\xeb\xb4ZI\xfd\xbc\xdax\xad\xe7\xdf\xfe\xf4\n0\xe7a#V\x91\xdf;\x9e\x16\a\u0097W\x8cl\xfc\x1d\v\xe7\x81\xd9\xdfwK(Ezq\x17}\xee\xc5\xde\xd8D\xfb\x024\xff\xa7\x15TF\xb1\x8d8\xfa\xe41=|g>\xd1\xca\x00\x9b\x11\x04\x12\x8fl@\xe5d\xbf\xd1*\xad\xb1c4\x0fǟ\xf4~\xf9\xe5\xe0\x8b]\xfeW\aol\xf9P\t\xff\xfcפXE\xf3<|\x96\x93\xc3\xff\x06\x00\x00\xff\xffS\x8a'\x11\"\x15\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcX\xcdr#\xb9\r\xbe\xeb)P\x9b\x83/Q{'\xbb\x95J\xe9\x96\xd5l\xaa\\\xc98S\xf6\xcc\\R9\xb0I\xb4\x9a1\x9b\xe4\x12\xa4l\xe5\xe9S ٭֏-9\x95I_l\x81\x00\b~\x00>\xa2{\xb1\\.\x17\xc2\xebo\x18H;\xbb\x02\xe15\xbeD\xb4\xfc\x8b\x9a\xa7?Q\xa3\xdd\xed\xf6\xc3\xe2I[\xb5\x82u\xa2\xe8\x86\a$\x97\x82ď\xd8i\xab\xa3vv1`\x14JD\xb1Z\x00\bk]\x14,&\xfe\t \x9d\x8d\xc1\x19\x83a\xb9A\xdb<\xa5\x16ۤ\x8d\u0090\x9d\x8f[o\u007fl~n~\\\x00Ȁ\xd9\xfc\x8b\x1e\x90\xa2\x18\xfc\nl2f\x01`ŀ+H\xde8\xa1\xa8\xe1\r\a\xb7\xc5\xd0HK\xca7\xdb\xe1Y\x04l\xa4\x1b\x16\xe4Q\xf2\xe6\x9b\xe0\x92_\xc1\x1b\x9a\xc5i\x8d\xb4\x9c\xf2k\xf6\x9f\x05FS\xfc\xebL\xf87M1/x\x93\x820S,YF\xdan\x92\x11a\x94.\x00H:\x8f+\xb8\xe7-\xbc\x90Ȳz\xe0\xbc\xe5\xb2\x1ei\xfbA\x18ߋ\x0fŏ\xecq\x10%\"\x00\xe7\xd1\xfe\xf9\xf3ݷ\x9f\x1e\x0f\xc4\x00\nI\x06\xedc\x06\xaf\x84We-\x82\x80-\x1a\fn\xe9M\xdah\v\xad\x90O\xc9O\xb6>8\x8f!\xea\xf1\xe0\xe5\x99\x15\xc2Lz\xb4\xd3\r\aS\xb4@q\x05 A\xecq<\x16\xaa\x1a?\xb8\x0eb\xaf\t\x02\xfa\x80\x84\xb6\xd4\x04\x8b\x85\x05\xd7\xfe\vel\xe0\x11\x03\x1b\x02\xf5.\x19ť\xb2\xc5\x10!\xa0t\x1b\xab\xff=y#\x88.ocDD\x8a\xa0m\xc4`\x85\x81\xad0\t\u007f\x0f\xc2*\x18\xc4\x0e\x02\xb2_Hv\xe6!\xabP\x03\x9f\\@жs+\xe8c\xf4\xb4\xba\xbd\xdd\xe88\x16\xb9tÐ\xac\x8e\xbb\xdb\\\xaf\xbaM\xd1\x05\xbaU\xb8EsKz\xb3\x14A\xf6:\xa2\x8c)\xe0\xad\xf0z\x99\x83\xb5\xb9ЛA\xfd.Զ\xa0\x9b\x03\xf0\xe2\x8e+\x80b\xd0v3[ȥ\xf6\x06\xca\\u\xa0\tD5-\xa7\u0603\xc9\"\xc6\xe3\xe1\xd7\xc7/0n]\x00/\xd8\xeeUi\x0f3C\xa4m\x87\xa1hv\xc1\r\xd9\vZ坶1\xff\x90F\xa3\x8d@\xa9\x1dt\xe4\xfc\xfd\x96\x90\"g\xa0\x81u\xeenh\x11\x92W\"\xa2j\xe0\xce\xc2Z\fhւ\xf0\xbb\x83\xcchҒ\xc1\xbb\x0e\xe691\x1d+\x17\x9cf\v#k\xbc\x92\x93G\x8f\x92S\x921\xcaL\xb8\a\x9eM\x0f,\xcfw\x18?\xa5\x15\x1f\xd0;\xd2х\xdd\xf1\xfaѮ\xbf\x1c\xa9\xb3\xe7\xadVH\xd5\x11\xe7y\\bܡs\xa1\x12P\x03_\tU\x16\f\xc9D\xed\r\x9e\x1a5'ۿ\x02\xe5>\xf6=1_\x13\xfa\xa4\x9d\xbb:\xa8\x02`\xd4\x03\xe6\u007fj@ς@\nc\xb8\xa2\xbe\xf4\b\x94+\xf6\x86\x8a\xa2&H\xe3Q\x1e\xad\xf0Ի8\xf9=\t\xa2sa\x101\xd3>.\xd9\xfe=G\xa4\xea\xfe\xee\xe3\x85\xd3=N\x8acQhŕ\xdai\f9P\x16\x8d\xde\n\x1b\"l\x9dI\x03\xbe\v\xf3\xd1\xc5\x03v\x18\xd0J\xbc2\xaeI\u007f\fώ\xb7Pf\xcb)2\x16\xd7x\x19\xe6r\x91Ԧ/\xb9(x\xb2\x9b\xd1\xf7t\xa1ݎ\x92\xf5\x03\xcb\xe0.N\xb9\x8a\xaerD\u07bc\xfa\xa5(b\"\xd0\xf6\x00\x9dw\xe1Q<\xad\x85\x95h.@\xf1u\xa6\n\xda*-\xf9\xf6\x18\x0f\xc7\x11ʲ\xe6\xec\xc61\xa3־y%\x9c\xd69\x83\xc2^E%\xf9\x9co\x90I\t\xed\xb1\xc21\x92J\b\x99y\x8b\x94oʪ\xd7\\\xc9-\xd2\r\xde\xe0\xe1\xf0\xf46D\xebS\x8b\xd3>\x15v\xcc_n\xd3b\u009d\xba\xb7\x9e\xfa\xb4\x18\xa3\x02ܢ\x05g\xa1\x13ڠ\x9a\xa6\xb6\v\xdd}&\x1e\xfa/\x1a\x9c\aF\xd1\x1a\\A\f\xe9]\xfd_\x93\xc0\xdc\xf5\xf7\xae\xbb\x84ށ\xf2\x01p\xcck\xae\xeb\x18\x81\x801\xec\xf2\xe9\x0e\xb0h\xe0!/\xb8\t\xdd:\x05\xb9\x16w\x80/\xdeY\xe6\x13a&_\x03\xca^XM\xc3i\x85\x8e\x90h\x1b\u007f\xfa\xc3+\xe7\xe5\xa1i\x83\xe1hu@\"\xb1\xb9\xc4*\x9f\x8aV\x19I\xaa\t\x88֥8k\xef\x1b\xaa\xb5\xfb\xae\x86\xb6\xf8\x123\x12\xd7\xd6\xec\xfd\x89\xc1\x88\\\x8bS\xd1\x16y외j\xe7\x17\x16ė\xf8ZF`\xfdP\xef\xcc\xe8\xa0\xd3&b\x80\xc3\x13N\xdc\xf1\xdck\xd9s+`f\xb3\x16;\x1e|f\xce9\x86\xd7\xd3\xf4=*\xd7\xf7\x82.\xa5\xf13\xeb\x9c\xe3\x1b\x1co\xa8s\x84\xc3\x0f\xda4\x9c:_\xc2=>\x9f\x91\xde\xd9\xcf\xc1m\x02\xd2i\xf3.\xc7\x1eGuf\xadl\xffk\b\xee\xb8N\xb3%\xf3o\xf2\u007f\xc9I;\xb7\x9e\xe9\xfc\x8d\xa5S\xd4ކ48\x89\xc4/v\xf7N]\u0096\x89\xed\xa3\x88⓰b\x83\x01\xacSX\n\xb0\x17\x04^˧\\h3\x94s\x05\xee\xf7`n\xd4\x04\xcfژل\r\x82\x80\x9c\xb3\xfcwf\xaa\xe7.\x8f\xfdܕ\\Σ\x91\xdc \xf6&\x8ez\xf3\x10\xc8\r\xfcj!\xc8Y\xd0q\n`\xef\xbf݁\xb0.\xf6\xf5T\xefjo_\v\xe1ReV5\xe8\x9d\x19\xef\x1e\x17\x85\x01\x9b\x86\x96\xfb\xb0\x83vǷ\xf7\xc1$\x95G\x99y-ϴ\xc5>\xfa\x88T\xc1\x95\xc2flk\x93+Mވ\xdd\x14d\x1e\xa1\xb9E\xf9J\xdb\xd3[\x1dR\xf9\xda\xcdK\xa7\x00\xbc~'\xf3\x93#\xf8\xe8\xec\x99\x12\x82C\xf2\xfe\xe3\xcfg5\xde\"\xf0\xbc\xceP\xfd»|\x9f\x1d^\x19s\xf8\xc9t\xb7v\xc9\xc6\v\x19~\x98\x14\x0f.\xca}\xc6\xf6\xc4I\xb9,\x85R<\x94\x89\x91\xa4k\xc5\x16\xea\xad2\x95\xb8N\xc0b|v\xe1\t4Q*\xe9b\xe9o\t\x13\xceނ\x12\xf1\x9bo\x10\xf2)\xfb\xb5\n\x14\xb6i\xb3\xe1\x9e\xf9\x1fި\x14E\x88\xd7^f\x8f\aʗf\xaf\xec\xfa\x8aw\xa4\x03\xa7\xff\xcf\x01\xeal\xa1\xf0\xa5\xa9\x03\xce>x,\x0fߘO\xac\xf2\xd1\xd4ls\x8a.\xf0\x94R$\xfb\xf1ZH\x89>\xa2\xba?\xfe\x8a\xf7\xc3\x0f\a\x1f\xea\xf2O\xe9\xac\xd2\xe5\xcb$\xfc㟋\xe2\x15շ\xf1\x8b\x1c\v\xff\x13\x00\x00\xff\xff\x81b\x1b\xc9\x13\x15\x00\x00"),
}

var CRDs = crds()

func crds() []*apiextv1.CustomResourceDefinition {
	apiextinstall.Install(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	var objs []*apiextv1.CustomResourceDefinition
	for _, crd := range rawCRDs {
		gzr, err := gzip.NewReader(bytes.NewReader(crd))
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(gzr)
		if err != nil {
			panic(err)
		}
		gzr.Close()

		obj, _, err := decode(bytes, nil, nil)
		if err != nil {
			panic(err)
		}
		objs = append(objs, obj.(*apiextv1.CustomResourceDefinition))
	}
	return objs
}
