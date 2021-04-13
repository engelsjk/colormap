package palette

// 'Rocket' palette licensed under a BSD 3-Clause "New" or "Revised" License found in LICENSES/seaborn.

type Rocket struct{}

func (r Rocket) LUT() [256][3]float32 {
	return [256][3]float32{
		{0.01060815, 0.01808215, 0.10018654},
		{0.01428972, 0.02048237, 0.10374486},
		{0.01831941, 0.0229766, 0.10738511},
		{0.02275049, 0.02554464, 0.11108639},
		{0.02759119, 0.02818316, 0.11483751},
		{0.03285175, 0.03088792, 0.11863035},
		{0.03853466, 0.03365771, 0.12245873},
		{0.04447016, 0.03648425, 0.12631831},
		{0.05032105, 0.03936808, 0.13020508},
		{0.05611171, 0.04224835, 0.13411624},
		{0.0618531, 0.04504866, 0.13804929},
		{0.06755457, 0.04778179, 0.14200206},
		{0.0732236, 0.05045047, 0.14597263},
		{0.0788708, 0.05305461, 0.14995981},
		{0.08450105, 0.05559631, 0.15396203},
		{0.09011319, 0.05808059, 0.15797687},
		{0.09572396, 0.06050127, 0.16200507},
		{0.10132312, 0.06286782, 0.16604287},
		{0.10692823, 0.06517224, 0.17009175},
		{0.1125315, 0.06742194, 0.17414848},
		{0.11813947, 0.06961499, 0.17821272},
		{0.12375803, 0.07174938, 0.18228425},
		{0.12938228, 0.07383015, 0.18636053},
		{0.13501631, 0.07585609, 0.19044109},
		{0.14066867, 0.0778224, 0.19452676},
		{0.14633406, 0.07973393, 0.1986151},
		{0.15201338, 0.08159108, 0.20270523},
		{0.15770877, 0.08339312, 0.20679668},
		{0.16342174, 0.0851396, 0.21088893},
		{0.16915387, 0.08682996, 0.21498104},
		{0.17489524, 0.08848235, 0.2190294},
		{0.18065495, 0.09009031, 0.22303512},
		{0.18643324, 0.09165431, 0.22699705},
		{0.19223028, 0.09317479, 0.23091409},
		{0.19804623, 0.09465217, 0.23478512},
		{0.20388117, 0.09608689, 0.23860907},
		{0.20973515, 0.09747934, 0.24238489},
		{0.21560818, 0.09882993, 0.24611154},
		{0.22150014, 0.10013944, 0.2497868},
		{0.22741085, 0.10140876, 0.25340813},
		{0.23334047, 0.10263737, 0.25697736},
		{0.23928891, 0.10382562, 0.2604936},
		{0.24525608, 0.10497384, 0.26395596},
		{0.25124182, 0.10608236, 0.26736359},
		{0.25724602, 0.10715148, 0.27071569},
		{0.26326851, 0.1081815, 0.27401148},
		{0.26930915, 0.1091727, 0.2772502},
		{0.27536766, 0.11012568, 0.28043021},
		{0.28144375, 0.11104133, 0.2835489},
		{0.2875374, 0.11191896, 0.28660853},
		{0.29364846, 0.11275876, 0.2896085},
		{0.29977678, 0.11356089, 0.29254823},
		{0.30592213, 0.11432553, 0.29542718},
		{0.31208435, 0.11505284, 0.29824485},
		{0.31826327, 0.1157429, 0.30100076},
		{0.32445869, 0.11639585, 0.30369448},
		{0.33067031, 0.11701189, 0.30632563},
		{0.33689808, 0.11759095, 0.3088938},
		{0.34314168, 0.11813362, 0.31139721},
		{0.34940101, 0.11863987, 0.3138355},
		{0.355676, 0.11910909, 0.31620996},
		{0.36196644, 0.1195413, 0.31852037},
		{0.36827206, 0.11993653, 0.32076656},
		{0.37459292, 0.12029443, 0.32294825},
		{0.38092887, 0.12061482, 0.32506528},
		{0.38727975, 0.12089756, 0.3271175},
		{0.39364518, 0.12114272, 0.32910494},
		{0.40002537, 0.12134964, 0.33102734},
		{0.40642019, 0.12151801, 0.33288464},
		{0.41282936, 0.12164769, 0.33467689},
		{0.41925278, 0.12173833, 0.33640407},
		{0.42569057, 0.12178916, 0.33806605},
		{0.43214263, 0.12179973, 0.33966284},
		{0.43860848, 0.12177004, 0.34119475},
		{0.44508855, 0.12169883, 0.34266151},
		{0.45158266, 0.12158557, 0.34406324},
		{0.45809049, 0.12142996, 0.34540024},
		{0.46461238, 0.12123063, 0.34667231},
		{0.47114798, 0.12098721, 0.34787978},
		{0.47769736, 0.12069864, 0.34902273},
		{0.48426077, 0.12036349, 0.35010104},
		{0.49083761, 0.11998161, 0.35111537},
		{0.49742847, 0.11955087, 0.35206533},
		{0.50403286, 0.11907081, 0.35295152},
		{0.51065109, 0.11853959, 0.35377385},
		{0.51728314, 0.1179558, 0.35453252},
		{0.52392883, 0.11731817, 0.35522789},
		{0.53058853, 0.11662445, 0.35585982},
		{0.53726173, 0.11587369, 0.35642903},
		{0.54394898, 0.11506307, 0.35693521},
		{0.5506426, 0.11420757, 0.35737863},
		{0.55734473, 0.11330456, 0.35775059},
		{0.56405586, 0.11235265, 0.35804813},
		{0.57077365, 0.11135597, 0.35827146},
		{0.5774991, 0.11031233, 0.35841679},
		{0.58422945, 0.10922707, 0.35848469},
		{0.59096382, 0.10810205, 0.35847347},
		{0.59770215, 0.10693774, 0.35838029},
		{0.60444226, 0.10573912, 0.35820487},
		{0.61118304, 0.10450943, 0.35794557},
		{0.61792306, 0.10325288, 0.35760108},
		{0.62466162, 0.10197244, 0.35716891},
		{0.63139686, 0.10067417, 0.35664819},
		{0.63812122, 0.09938212, 0.35603757},
		{0.64483795, 0.0980891, 0.35533555},
		{0.65154562, 0.09680192, 0.35454107},
		{0.65824241, 0.09552918, 0.3536529},
		{0.66492652, 0.09428017, 0.3526697},
		{0.67159578, 0.09306598, 0.35159077},
		{0.67824099, 0.09192342, 0.3504148},
		{0.684863, 0.09085633, 0.34914061},
		{0.69146268, 0.0898675, 0.34776864},
		{0.69803757, 0.08897226, 0.3462986},
		{0.70457834, 0.0882129, 0.34473046},
		{0.71108138, 0.08761223, 0.3430635},
		{0.7175507, 0.08716212, 0.34129974},
		{0.72398193, 0.08688725, 0.33943958},
		{0.73035829, 0.0868623, 0.33748452},
		{0.73669146, 0.08704683, 0.33543669},
		{0.74297501, 0.08747196, 0.33329799},
		{0.74919318, 0.08820542, 0.33107204},
		{0.75535825, 0.08919792, 0.32876184},
		{0.76145589, 0.09050716, 0.32637117},
		{0.76748424, 0.09213602, 0.32390525},
		{0.77344838, 0.09405684, 0.32136808},
		{0.77932641, 0.09634794, 0.31876642},
		{0.78513609, 0.09892473, 0.31610488},
		{0.79085854, 0.10184672, 0.313391},
		{0.7965014, 0.10506637, 0.31063031},
		{0.80205987, 0.10858333, 0.30783},
		{0.80752799, 0.11239964, 0.30499738},
		{0.81291606, 0.11645784, 0.30213802},
		{0.81820481, 0.12080606, 0.29926105},
		{0.82341472, 0.12535343, 0.2963705},
		{0.82852822, 0.13014118, 0.29347474},
		{0.83355779, 0.13511035, 0.29057852},
		{0.83850183, 0.14025098, 0.2876878},
		{0.84335441, 0.14556683, 0.28480819},
		{0.84813096, 0.15099892, 0.281943},
		{0.85281737, 0.15657772, 0.27909826},
		{0.85742602, 0.1622583, 0.27627462},
		{0.86196552, 0.16801239, 0.27346473},
		{0.86641628, 0.17387796, 0.27070818},
		{0.87079129, 0.17982114, 0.26797378},
		{0.87507281, 0.18587368, 0.26529697},
		{0.87925878, 0.19203259, 0.26268136},
		{0.8833417, 0.19830556, 0.26014181},
		{0.88731387, 0.20469941, 0.25769539},
		{0.89116859, 0.21121788, 0.2553592},
		{0.89490337, 0.21785614, 0.25314362},
		{0.8985026, 0.22463251, 0.25108745},
		{0.90197527, 0.23152063, 0.24918223},
		{0.90530097, 0.23854541, 0.24748098},
		{0.90848638, 0.24568473, 0.24598324},
		{0.911533, 0.25292623, 0.24470258},
		{0.9144225, 0.26028902, 0.24369359},
		{0.91717106, 0.26773821, 0.24294137},
		{0.91978131, 0.27526191, 0.24245973},
		{0.92223947, 0.28287251, 0.24229568},
		{0.92456587, 0.29053388, 0.24242622},
		{0.92676657, 0.29823282, 0.24285536},
		{0.92882964, 0.30598085, 0.24362274},
		{0.93078135, 0.31373977, 0.24468803},
		{0.93262051, 0.3215093, 0.24606461},
		{0.93435067, 0.32928362, 0.24775328},
		{0.93599076, 0.33703942, 0.24972157},
		{0.93752831, 0.34479177, 0.25199928},
		{0.93899289, 0.35250734, 0.25452808},
		{0.94036561, 0.36020899, 0.25734661},
		{0.94167588, 0.36786594, 0.2603949},
		{0.94291042, 0.37549479, 0.26369821},
		{0.94408513, 0.3830811, 0.26722004},
		{0.94520419, 0.39062329, 0.27094924},
		{0.94625977, 0.39813168, 0.27489742},
		{0.94727016, 0.4055909, 0.27902322},
		{0.94823505, 0.41300424, 0.28332283},
		{0.94914549, 0.42038251, 0.28780969},
		{0.95001704, 0.42771398, 0.29244728},
		{0.95085121, 0.43500005, 0.29722817},
		{0.95165009, 0.44224144, 0.30214494},
		{0.9524044, 0.44944853, 0.3072105},
		{0.95312556, 0.45661389, 0.31239776},
		{0.95381595, 0.46373781, 0.31769923},
		{0.95447591, 0.47082238, 0.32310953},
		{0.95510255, 0.47787236, 0.32862553},
		{0.95569679, 0.48489115, 0.33421404},
		{0.95626788, 0.49187351, 0.33985601},
		{0.95681685, 0.49882008, 0.34555431},
		{0.9573439, 0.50573243, 0.35130912},
		{0.95784842, 0.51261283, 0.35711942},
		{0.95833051, 0.51946267, 0.36298589},
		{0.95879054, 0.52628305, 0.36890904},
		{0.95922872, 0.53307513, 0.3748895},
		{0.95964538, 0.53983991, 0.38092784},
		{0.96004345, 0.54657593, 0.3870292},
		{0.96042097, 0.55328624, 0.39319057},
		{0.96077819, 0.55997184, 0.39941173},
		{0.9611152, 0.5666337, 0.40569343},
		{0.96143273, 0.57327231, 0.41203603},
		{0.96173392, 0.57988594, 0.41844491},
		{0.96201757, 0.58647675, 0.42491751},
		{0.96228344, 0.59304598, 0.43145271},
		{0.96253168, 0.5995944, 0.43805131},
		{0.96276513, 0.60612062, 0.44471698},
		{0.96298491, 0.6126247, 0.45145074},
		{0.96318967, 0.61910879, 0.45824902},
		{0.96337949, 0.6255736, 0.46511271},
		{0.96355923, 0.63201624, 0.47204746},
		{0.96372785, 0.63843852, 0.47905028},
		{0.96388426, 0.64484214, 0.4861196},
		{0.96403203, 0.65122535, 0.4932578},
		{0.96417332, 0.65758729, 0.50046894},
		{0.9643063, 0.66393045, 0.5077467},
		{0.96443322, 0.67025402, 0.51509334},
		{0.96455845, 0.67655564, 0.52251447},
		{0.96467922, 0.68283846, 0.53000231},
		{0.96479861, 0.68910113, 0.53756026},
		{0.96492035, 0.69534192, 0.5451917},
		{0.96504223, 0.7015636, 0.5528892},
		{0.96516917, 0.70776351, 0.5606593},
		{0.96530224, 0.71394212, 0.56849894},
		{0.96544032, 0.72010124, 0.57640375},
		{0.96559206, 0.72623592, 0.58438387},
		{0.96575293, 0.73235058, 0.59242739},
		{0.96592829, 0.73844258, 0.60053991},
		{0.96612013, 0.74451182, 0.60871954},
		{0.96632832, 0.75055966, 0.61696136},
		{0.96656022, 0.75658231, 0.62527295},
		{0.96681185, 0.76258381, 0.63364277},
		{0.96709183, 0.76855969, 0.64207921},
		{0.96739773, 0.77451297, 0.65057302},
		{0.96773482, 0.78044149, 0.65912731},
		{0.96810471, 0.78634563, 0.66773889},
		{0.96850919, 0.79222565, 0.6764046},
		{0.96893132, 0.79809112, 0.68512266},
		{0.96935926, 0.80395415, 0.69383201},
		{0.9698028, 0.80981139, 0.70252255},
		{0.97025511, 0.81566605, 0.71120296},
		{0.97071849, 0.82151775, 0.71987163},
		{0.97120159, 0.82736371, 0.72851999},
		{0.97169389, 0.83320847, 0.73716071},
		{0.97220061, 0.83905052, 0.74578903},
		{0.97272597, 0.84488881, 0.75440141},
		{0.97327085, 0.85072354, 0.76299805},
		{0.97383206, 0.85655639, 0.77158353},
		{0.97441222, 0.86238689, 0.78015619},
		{0.97501782, 0.86821321, 0.78871034},
		{0.97564391, 0.87403763, 0.79725261},
		{0.97628674, 0.87986189, 0.8057883},
		{0.97696114, 0.88568129, 0.81430324},
		{0.97765722, 0.89149971, 0.82280948},
		{0.97837585, 0.89731727, 0.83130786},
		{0.97912374, 0.90313207, 0.83979337},
		{0.979891, 0.90894778, 0.84827858},
		{0.98067764, 0.91476465, 0.85676611},
		{0.98137749, 0.92061729, 0.86536915},
	}
}

func (r Rocket) Lookup(x uint8) [3]float32 {
	return r.LUT()[x]
}