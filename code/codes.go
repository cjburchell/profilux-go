package code

const (
	// Codes
	// Codes welche nicht im EEPROM gespeichert werden sind ab= 1000

	// Codes für Kommunikation
	// Codes fuer allgem. Werte
	// ReSharper disable UnusedMember.Global
	// ReSharper disable InconsistentNaming

	// Codes für Kommunikation
	// Codes fuer allgem. Werte

	// The softwareversion.

	SOFTWAREVERSION = 0

	// The softwaredate.
	SOFTWAREDATE = 1

	// The productid.
	PRODUCTID = 2

	// The loaddefaults.
	LOADDEFAULTS = 3

	// xxx = 4 //frei!

	// The address.
	ADDRESS = 5

	// The serialnumber.
	SERIALNUMBER = 6

	// danach Platz lassen wenn Wert ins EEPROM geschrieben werden muss (Serialnumber = ULONG)
	// 7 ist frei

	// The memorystate.

	MEMORYSTATE = 8

	// The profiluxview.

	PROFILUXVIEW = 9

	// The sensorpar a 1_ ca l 1 value.

	SENSORPARA1_CAL1VALUE = 10

	// The sensorpar a 1_ ca l 2 value.

	SENSORPARA1_CAL2VALUE = 11

	// The sensorpar a 1_ ca l 1 adc.

	SENSORPARA1_CAL1ADC = 12

	// The sensorpar a 1_ ca l 2 adc.

	SENSORPARA1_CAL2ADC = 13

	// The sensorpar a 1_ alarmmode.

	SENSORPARA1_ALARMMODE = 14

	// The sensorpar a 1_ desvalue.

	SENSORPARA1_DESVALUE = 15

	// The sensorpar a 1_ hyst.

	SENSORPARA1_HYST = 16

	// The sensorpar a 1_ ncenabled.

	SENSORPARA1_NCENABLED = 17

	// The sensorpar a 1_ ncvalue.

	SENSORPARA1_NCVALUE = 18

	// The sensorpar a 1_ starttime.

	SENSORPARA1_STARTTIME = 19

	// The sensorpar a 1_ endtime.

	SENSORPARA1_ENDTIME = 20

	// The sensorpar a 1_ enabled.

	SENSORPARA1_PROPS = 21

	// The sensorpar a 1_ par a 1.

	SENSORPARA1_PARA1 = 22

	// The sensorpar a 1_ par a 2.

	SENSORPARA1_PARA2 = 23

	// The sensorpar a 1_ alarmdeviation.

	SENSORPARA1_ALARMDEVIATION = 24

	// The sensorpar a 1_ sensortype.

	SENSORPARA1_SENSORTYPE = 25

	// The sensorpar a 1_ par a 3.

	SENSORPARA1_PARA3 = 26

	// The sensorpar a 1_ displaymode.

	SENSORPARA1_DISPLAYMODE = 27

	// The sensorpar a 1_ pulsewidth.

	SENSORPARA1_PULSEWIDTH = 28

	// The sensorpar a 1_ pausewidth.

	SENSORPARA1_PAUSEWIDTH = 29

	// The sensorpar a 1_ controllingmode.

	SENSORPARA1_CONTROLLINGMODE = 30

	// The sensorpar a 1_ par a 4.

	SENSORPARA1_PARA4 = 31

	// The sensorpar a 1_ par a 5.

	SENSORPARA1_PARA5 = 32

	// The sensorpar a 1_ signalfilter.

	SENSORPARA1_SIGNALFILTER = 33

	// The sensorpar a 8_ signalfilter.

	SENSORPARA8_SIGNALFILTER = 201

	// Codes Dimmkanäle (28 pro Dimmkanal 8 Dimmkanäle)

	// The illuminatio n 1_ simulationmask.

	ILLUMINATION1_SIMULATIONMASK = 202

	// The illuminatio n 1_ burninduration.

	ILLUMINATION1_BURNINDURATION = 203

	// The illuminatio n 1_ props.

	ILLUMINATION1_PROPS = 204

	// The illuminatio n 1_ tim e 1.

	ILLUMINATION1_TIME1 = 205

	// The illuminatio n 1_ brightnes s 1.

	ILLUMINATION1_BRIGHTNESS1 = 206

	// The illuminatio n 1_ tim e 2.

	ILLUMINATION1_TIME2 = 207

	// The illuminatio n 1_ brightnes s 2.

	ILLUMINATION1_BRIGHTNESS2 = 208

	// The illuminatio n 1_ tim e 3.

	ILLUMINATION1_TIME3 = 209

	// The illuminatio n 1_ brightnes s 3.

	ILLUMINATION1_BRIGHTNESS3 = 210

	// The illuminatio n 1_ tim e 4.

	ILLUMINATION1_TIME4 = 211

	// The illuminatio n 1_ brightnes s 4.

	ILLUMINATION1_BRIGHTNESS4 = 212

	// The illuminatio n 1_ tim e 5.

	ILLUMINATION1_TIME5 = 213

	// The illuminatio n 1_ brightnes s 5.

	ILLUMINATION1_BRIGHTNESS5 = 214

	// The illuminatio n 1_ tim e 6.

	ILLUMINATION1_TIME6 = 215

	// The illuminatio n 1_ brightnes s 6.

	ILLUMINATION1_BRIGHTNESS6 = 216

	// The illuminatio n 1_ tim e 7.

	ILLUMINATION1_TIME7 = 217

	// The illuminatio n 1_ brightnes s 7.

	ILLUMINATION1_BRIGHTNESS7 = 218

	// The illuminatio n 1_ tim e 8.

	ILLUMINATION1_TIME8 = 219

	// The illuminatio n 1_ brightnes s 8.

	ILLUMINATION1_BRIGHTNESS8 = 220

	// The illuminatio n 1_ tim e 9.

	ILLUMINATION1_TIME9 = 221

	// The illuminatio n 1_ brightnes s 9.

	ILLUMINATION1_BRIGHTNESS9 = 222

	// The illuminatio n 1_ tim e 10.

	ILLUMINATION1_TIME10 = 223

	// The illuminatio n 1_ brightnes s 10.

	ILLUMINATION1_BRIGHTNESS10 = 224

	// The illuminatio n 1_ tim e 11.

	ILLUMINATION1_TIME11 = 225

	// The illuminatio n 1_ brightnes s 11.

	ILLUMINATION1_BRIGHTNESS11 = 226

	// The illuminatio n 1_ tim e 12.

	ILLUMINATION1_TIME12 = 227

	// The illuminatio n 1_ brightnes s 12.

	ILLUMINATION1_BRIGHTNESS12 = 228

	// The illuminatio n 1_ reserved.

	ILLUMINATION1_RESERVED = 229

	// The illuminatio n 8_ reserved.

	ILLUMINATION8_RESERVED = 425

	// Codes Erinnerung (12 pro Erinnerung 4 Erinnerungen)

	// The me m 1_ nextmem.

	MEM1_NEXTMEM = 426

	// The me m 1_ repeat.

	MEM1_REPEAT = 427

	// The me m 1_ days.

	MEM1_DAYS = 428

	// The me m 1_ tex t 01.

	MEM1_TEXT01 = 429

	// The me m 1_ tex t 23.

	MEM1_TEXT23 = 430

	// The me m 1_ tex t 45.

	MEM1_TEXT45 = 431

	// The me m 1_ tex t 67.

	MEM1_TEXT67 = 432

	// The me m 1_ tex t 89.

	MEM1_TEXT89 = 433

	// The me m 1_ tex t 1011.

	MEM1_TEXT1011 = 434

	// The me m 1_ tex t 1213.

	MEM1_TEXT1213 = 435

	// The me m 1_ tex t 1415.

	MEM1_TEXT1415 = 436

	// The me m 1_ reserved.

	MEM1_RESERVED = 437

	// The me m 4_ reserved.

	MEM4_RESERVED = 473

	// Codes Zeitschaltuhren (21 pro Zeitschaltuhr 12 Zeitschaltuhren)
	// Zeitschaltuhr 1

	// The time r 1_ props.

	TIMER1_PROPS = 474

	// The time r 1_ rateperdosing.

	TIMER1_RATEPERDOSING = 475

	// The time r 1_ dayoffset.

	TIMER1_DAYOFFSET = 476

	// The time r 1_ flowrate.

	TIMER1_FLOWRATE = 477

	// The time r 1_ do w_ int.

	TIMER1_DOW_INT = 478

	// The time r 1_ star t 1.

	TIMER1_START1 = 479

	// The time r 1_ en d 1.

	TIMER1_END1 = 480

	// The time r 1_ star t 2.

	TIMER1_START2 = 481

	// The time r 1_ en d 2.

	TIMER1_END2 = 482

	// The time r 1_ star t 3.

	TIMER1_START3 = 483

	// The time r 1_ en d 3.

	TIMER1_END3 = 484

	// The time r 1_ star t 4.

	TIMER1_START4 = 485

	// The time r 1_ en d 4.

	TIMER1_END4 = 486

	// The time r 1_ star t 5.

	TIMER1_START5 = 487

	// The time r 1_ en d 5.

	TIMER1_END5 = 488

	// The time r 1_ star t 6.

	TIMER1_START6 = 489

	// The time r 1_ en d 6.

	TIMER1_END6 = 490

	// The time r 1_ star t 7.

	TIMER1_START7 = 491

	// The time r 1_ en d 7.

	TIMER1_END7 = 492

	// The time r 1_ star t 8.

	TIMER1_START8 = 493

	// The time r 1_ en d 8.

	TIMER1_END8 = 494

	// The time r 12_ en d 8.

	TIMER12_END8 = 725

	// Konfiguration 1-10V-Schnittstellen

	// The l 1 t o 10 vin t 1_ umin.

	L1TO10VINT1_UMIN = 726

	// The l 1 t o 10 vin t 1_ umax.

	L1TO10VINT1_UMAX = 727

	// The l 1 t o 10 vin t 1_ function.

	L1TO10VINT1_FUNCTION = 728

	// ...

	// The l 1 t o 10 vin t 10_ function.

	L1TO10VINT10_FUNCTION = 755

	// Codes Steckdosen

	// The switchplu g 1_ function.

	SWITCHPLUG1_FUNCTION = 756

	// ...

	// The switchplu g 24_ function.

	SWITCHPLUG24_FUNCTION = 779

	// The digitalpowerbaravailable.

	DIGITALPOWERBARAVAILABLE = 780

	// 781 - 803 frei!!!

	// Wartung

	// The maintenanc e_ spselmas k 1.

	MAINTENANCE_SPSELMASK1 = 804

	// The maintenanc e_ spselmas k 2.

	MAINTENANCE_SPSELMASK2 = 805

	// The maintenanc e_ spstatemas k 1.

	MAINTENANCE_SPSTATEMASK1 = 806

	// The maintenanc e_ spstatemas k 2.

	MAINTENANCE_SPSTATEMASK2 = 807

	// The maintenanc e_ onetotenpercen t 1.

	MAINTENANCE_ONETOTENPERCENT1 = 808

	// The maintenanc e_ onetotenpercen t 2.

	MAINTENANCE_ONETOTENPERCENT2 = 809

	// The maintenanc e_ onetotenpercen t 3.

	MAINTENANCE_ONETOTENPERCENT3 = 810

	// The maintenanc e_ onetotenpercen t 4.

	MAINTENANCE_ONETOTENPERCENT4 = 811

	// The maintenanc e_ onetotenpercen t 5.

	MAINTENANCE_ONETOTENPERCENT5 = 812

	// The maintenanc e_ onetotenpercen t 6.

	MAINTENANCE_ONETOTENPERCENT6 = 813

	// The maintenanc e_ onetotenpercen t 7.

	MAINTENANCE_ONETOTENPERCENT7 = 814

	// The maintenanc e_ onetotenpercen t 8.

	MAINTENANCE_ONETOTENPERCENT8 = 815

	// The maintenanc e_ onetotenpercen t 9.

	MAINTENANCE_ONETOTENPERCENT9 = 816

	// The maintenanc e_ onetotenpercen t 10.

	MAINTENANCE_ONETOTENPERCENT10 = 817

	// The maintenanc e_ onetotenselmask.

	MAINTENANCE_ONETOTENSELMASK = 818

	// The maintenanc e_ timeout.

	MAINTENANCE_TIMEOUT = 816

	// The pin.

	PIN = 820

	// The time.

	TIME = 821

	// The date.

	DATE = 822

	// The dc f_ active.

	DCF_ACTIVE = 823

	// The mesmes z_ change.

	MESMESZ_CHANGE = 824

	// The correctionperday.

	CORRECTIONPERDAY = 825

	// The changetime.

	CHANGETIME = 826

	// The userinterface.

	USERINTERFACE = 827

	// The lo c_ longitude.

	LOC_LONGITUDE = 828

	// The lo c_ latitude.

	LOC_LATITUDE = 829

	// The alarmbeepmode.

	ALARMBEEPMODE = 830

	// The alarmbeepstarttime.

	ALARMBEEPSTARTTIME = 831

	// The alarmbeependtime.

	ALARMBEEPENDTIME = 832

	// Bank 1

	// The ma c 0.

	MAC0 = 1820

	// The ma c 1.

	MAC1 = 1821

	// The ma c 2.

	MAC2 = 1822

	// The ma c 3.

	MAC3 = 1823

	// The ma c 4.

	MAC4 = 1824

	// The ma c 5.

	MAC5 = 1825

	// The pabaddress.

	PABADDRESS = 1826

	// The localip.

	LOCALIP = 1827

	// The defgw.

	DEFGW = 1828

	// The netmask.

	NETMASK = 1829

	// Web Server Port

	WEBSERVERPORT = 1830

	// Gewitter

	// The thunderstor m_ intensity.

	THUNDERSTORM_INTENSITY = 833

	// The thunderstor m_ darkening.

	THUNDERSTORM_DARKENING = 834

	// The thunderstor m_ count.

	THUNDERSTORM_COUNT = 835

	// The thunderstor m_ star t 1.

	THUNDERSTORM_START1 = 836

	// The thunderstor m_ duratio n 1.

	THUNDERSTORM_DURATION1 = 837

	// The thunderstor m_ star t 2.

	THUNDERSTORM_START2 = 838

	// The thunderstor m_ duratio n 2.

	THUNDERSTORM_DURATION2 = 839

	// The thunderstor m_ star t 3.

	THUNDERSTORM_START3 = 840

	// The thunderstor m_ duratio n 3.

	THUNDERSTORM_DURATION3 = 841

	// The thunderstor m_ star t 4.

	THUNDERSTORM_START4 = 842

	// The thunderstor m_ duratio n 4.

	THUNDERSTORM_DURATION4 = 843

	// The thunderstor m_ dow.

	THUNDERSTORM_DOW = 844

	// The thunderstor m_ rndduration.

	THUNDERSTORM_RNDDURATION = 845

	// The thunderstor m_ rndminwait.

	THUNDERSTORM_RNDMINWAIT = 846

	// The thunderstor m_ rndmaxwait.

	THUNDERSTORM_RNDMAXWAIT = 847

	// temperaturabhängige Lichtreduzierung

	// The tempdeplightre d_ illuminationchannels.

	TEMPDEPLIGHTRED_ILLUMINATIONCHANNELS = 848

	// The tempdeplightre d_ deltatmin.

	TEMPDEPLIGHTRED_DELTATMIN = 849

	// The tempdeplightre d_ deltatmax.

	TEMPDEPLIGHTRED_DELTATMAX = 850

	// The tempdeplightre d_ sensorindex.

	TEMPDEPLIGHTRED_SENSORINDEX = 851

	// The tempdeplightre d_ deltatshutoff.

	TEMPDEPLIGHTRED_DELTATSHUTOFF = 852

	// Anzeigeeinstellungen

	// The displa y_ changetime.

	DISPLAY_CHANGETIME = 853

	// The displa y_ showmas k 1.

	DISPLAY_SHOWMASK1 = 854

	// The displa y_ showmas k 2.

	DISPLAY_SHOWMASK2 = 855

	// The displa y_ datetimemode.

	DISPLAY_DATETIMEMODE = 856

	// Bank 1

	// The displa y_ bright.

	DISPLAY_BRIGHT = 1853

	// The displa y_ dark.

	DISPLAY_DARK = 1854

	// The displa y_ darkstarttime.

	DISPLAY_DARKSTARTTIME = 1855

	// The displa y_ darkendtime.

	DISPLAY_DARKENDTIME = 1856

	// Messwerterfassung (weitere unten)

	// The measuremen t_ sampleperiod.

	MEASUREMENT_SAMPLEPERIOD = 857

	// The measuremen t_ samplesourcemask.

	MEASUREMENT_SAMPLESOURCEMASK = 858

	// The measuremen t_ maxmemorysize.

	MEASUREMENT_MAXMEMORYSIZE = 859

	// Stroemung

	// The currentcontro l_ nightdecactive.

	CURRENTCONTROL_NIGHTDECACTIVE = 860

	// The currentcontro l_ nightdecstart.

	CURRENTCONTROL_NIGHTDECSTART = 861

	// The currentcontro l_ nightdecend.

	CURRENTCONTROL_NIGHTDECEND = 862

	// 2 frei

	// The currentcontro l_ grou p 1 pumpcount.

	CURRENTCONTROL_GROUP1PUMPCOUNT = 865

	// The currentcontro l_ grou p 1 mode.

	CURRENTCONTROL_GROUP1MODE = 866

	// The currentcontro l_ grou p 1 minduration.

	CURRENTCONTROL_GROUP1MINDURATION = 867

	// The currentcontro l_ grou p 1 maxduration.

	CURRENTCONTROL_GROUP1MAXDURATION = 868

	// The currentcontro l_ grou p 1 minwaveduration.

	CURRENTCONTROL_GROUP1MINWAVEDURATION = 869

	// The currentcontro l_ grou p 1 waveform.

	CURRENTCONTROL_GROUP1WAVEFORM = 870

	// The currentcontro l_ grou p 1 rndreduction.

	CURRENTCONTROL_GROUP1RNDREDUCTION = 871

	// The currentcontro l_ grou p 1 maxwaveduration.

	CURRENTCONTROL_GROUP1MAXWAVEDURATION = 872

	// 2 frei

	// The currentcontro l_ grou p 2 pumpcount.

	CURRENTCONTROL_GROUP2PUMPCOUNT = 875

	// The currentcontro l_ grou p 2 mode.

	CURRENTCONTROL_GROUP2MODE = 876

	// The currentcontro l_ grou p 2 minduration.

	CURRENTCONTROL_GROUP2MINDURATION = 877

	// The currentcontro l_ grou p 2 maxduration.

	CURRENTCONTROL_GROUP2MAXDURATION = 878

	// The currentcontro l_ grou p 2 minwaveduration.

	CURRENTCONTROL_GROUP2MINWAVEDURATION = 879

	// The currentcontro l_ grou p 2 waveform.

	CURRENTCONTROL_GROUP2WAVEFORM = 880

	// The currentcontro l_ grou p 2 rndreduction.

	CURRENTCONTROL_GROUP2RNDREDUCTION = 881

	// The currentcontro l_ grou p 2 maxwaveduration.

	CURRENTCONTROL_GROUP2MAXWAVEDURATION = 882

	// 2 frei

	// The currentcontro l_ pum p 1 min.

	CURRENTCONTROL_PUMP1MIN = 885

	// The currentcontro l_ pum p 1 max.

	CURRENTCONTROL_PUMP1MAX = 886

	// The currentcontro l_ pum p 1 nightvalue.

	CURRENTCONTROL_PUMP1NIGHTVALUE = 887

	// The currentcontro l_ pum p 1 thundersvalue.

	CURRENTCONTROL_PUMP1THUNDERSVALUE = 888

	// The currentcontro l_ pum p 1 feedpausemode.

	CURRENTCONTROL_PUMP1FEEDPAUSEMODE = 889

	// 2 frei

	// The currentcontro l_ pum p 2 min.

	CURRENTCONTROL_PUMP2MIN = 892

	// The currentcontro l_ pum p 2 max.

	CURRENTCONTROL_PUMP2MAX = 893

	// The currentcontro l_ pum p 2 nightvalue.

	CURRENTCONTROL_PUMP2NIGHTVALUE = 894

	// The currentcontro l_ pum p 2 thundersvalue.

	CURRENTCONTROL_PUMP2THUNDERSVALUE = 895

	// The currentcontro l_ pum p 2 feedpausemode.

	CURRENTCONTROL_PUMP2FEEDPAUSEMODE = 896

	// 2 frei

	// The currentcontro l_ pum p 3 min.

	CURRENTCONTROL_PUMP3MIN = 899

	// The currentcontro l_ pum p 3 max.

	CURRENTCONTROL_PUMP3MAX = 900

	// The currentcontro l_ pum p 3 nightvalue.

	CURRENTCONTROL_PUMP3NIGHTVALUE = 901

	// The currentcontro l_ pum p 3 thundersvalue.

	CURRENTCONTROL_PUMP3THUNDERSVALUE = 902

	// The currentcontro l_ pum p 3 feedpausemode.

	CURRENTCONTROL_PUMP3FEEDPAUSEMODE = 903

	// 2 frei

	// The currentcontro l_ pum p 4 min.

	CURRENTCONTROL_PUMP4MIN = 906

	// The currentcontro l_ pum p 4 max.

	CURRENTCONTROL_PUMP4MAX = 907

	// The currentcontro l_ pum p 4 nightvalue.

	CURRENTCONTROL_PUMP4NIGHTVALUE = 908

	// The currentcontro l_ pum p 4 thundersvalue.

	CURRENTCONTROL_PUMP4THUNDERSVALUE = 909

	// The currentcontro l_ pum p 4 feedpausemode.

	CURRENTCONTROL_PUMP4FEEDPAUSEMODE = 910

	// Variable Beleuchtung

	// The variableilluminatio n 1_ l.

	VARIABLEILLUMINATION1_L = 911

	// The variableilluminatio n 1_ h.

	VARIABLEILLUMINATION1_H = 912

	// The variableilluminatio n 2_ l.

	VARIABLEILLUMINATION2_L = 913

	// The variableilluminatio n 2_ h.

	VARIABLEILLUMINATION2_H = 914

	// The variableilluminatio n 3_ l.

	VARIABLEILLUMINATION3_L = 915

	// The variableilluminatio n 3_ h.

	VARIABLEILLUMINATION3_H = 916

	// The variableilluminatio n 4_ l.

	VARIABLEILLUMINATION4_L = 917

	// The variableilluminatio n 4_ h.

	VARIABLEILLUMINATION4_H = 918

	// digitale Eingänge

	// The digitalinpu t 1_ function.

	DIGITALINPUT1_FUNCTION = 919

	// The digitalinpu t 2_ function.

	DIGITALINPUT2_FUNCTION = 920

	// The digitalinpu t 3_ function.

	DIGITALINPUT3_FUNCTION = 921

	// The digitalinpu t 4_ function.

	DIGITALINPUT4_FUNCTION = 922

	// COM

	// The co m 1_ baudrate.

	COM1_BAUDRATE = 923

	// The co m 2_ baudrate.

	COM2_BAUDRATE = 924

	// The co m 3_ baudrate.

	COM3_BAUDRATE = 925

	// The co m 4_ baudrate.

	COM4_BAUDRATE = 926

	// Niveau-Regelung

	// The leve l 1_ props.

	LEVEL1_PROPS = 927

	// The leve l 1_ reactduration.

	LEVEL1_SOURCES = 928

	// The leve l 1_ maxduration.

	LEVEL1_MAXDURATION = 929

	// The leve l 2_ props.

	LEVEL2_PROPS = 930

	// The leve l 2_ reactduration.

	LEVEL2_REACTDURATION = 931

	// The leve l 2_ maxduration.

	LEVEL2_MAXDURATION = 932

	// The leve l 3_ props.

	LEVEL3_PROPS = 933

	// The leve l 3_ reactduration.

	LEVEL3_REACTDURATION = 934

	// The leve l 3_ maxduration.

	LEVEL3_MAXDURATION = 935

	// The feedpaus e_ duration.

	FEEDPAUSE_DURATION = 936

	// The feedpaus e_ mode.

	FEEDPAUSE_MODE = 937

	// The rainingday s_ dow.

	RAININGDAYS_DOW = 938

	// The rainingday s_ darkening.

	RAININGDAYS_DARKENING = 939

	// Wolkensimulation

	// The clou d_ probability.

	CLOUD_PROBABILITY = 940

	// The clou d_ minduration.

	CLOUD_MINDURATION = 941

	// The clou d_ maxduration.

	CLOUD_MAXDURATION = 942

	// The clou d_ maxdarkening.

	CLOUD_MAXDARKENING = 943

	// Mondphasensimulation

	// The moo n_ start.

	MOON_START = 944

	// The moo n_ end.

	MOON_END = 945

	// The dal i_ mindim.

	DALI_MINDIM = 946

	// The aquailluminatio n_ available.

	AQUAILLUMINATION_AVAILABLE = 947

	// The aquailluminatio n_ whitechannel.

	AQUAILLUMINATION_WHITECHANNEL = 948

	// The aquailluminatio n_ bluechannel.

	AQUAILLUMINATION_BLUECHANNEL = 949

	// Codes für progr. Logik

	// The proglogi c 1_ inpu t 1.

	PROGLOGIC1_INPUT1 = 950

	// The proglogi c 1_ inpu t 2.

	PROGLOGIC1_INPUT2 = 951

	// The proglogi c 1_ function.

	PROGLOGIC1_FUNCTION = 952

	// The proglogi c 1_ time.

	PROGLOGIC1_TIME = 953

	// The proglogi c 8_ time.

	PROGLOGIC8_TIME = 981

	// The ca l 1 ad c_ tem p 1.

	CAL1ADC_TEMP1 = 982

	// The ca l 2 ad c_ tem p 1.

	CAL2ADC_TEMP1 = 983

	// ACHTUNG: keine Codes ab inkl. 984 die im EEPROM gespeichert werden!!!
	// da kommen feste EEPROM-Adressen!!!

	// from here only volatile data (not stored in EEPROM)

	// The sensorpar a 1_ actvalue.

	SENSORPARA1_ACTVALUE = 10000

	// The sensorpar a 1_ actadc.

	SENSORPARA1_ACTADC = 10001

	// The sensorpar a 1_ ohm.

	SENSORPARA1_OHM = 10002

	// The sensorpar a 1_ actstate.

	SENSORPARA1_ACTSTATE = 10003

	// The sensorpar a 1_ actanalogout.

	SENSORPARA1_ACTANALOGOUT = 10004

	// The sensorpar a 1_ actre s 1.

	SENSORPARA1_ACTRES1 = 10005

	// The sensorpar a 1_ actre s 2.

	SENSORPARA1_ACTRES2 = 10006

	// The sensorpar a 1_ actre s 3.

	SENSORPARA1_ACTRES3 = 10007

	// ...

	// The sensorpar a 8_ actre s 3.

	SENSORPARA8_ACTRES3 = 10063

	// PTC

	// The progra m_ ptc.

	PROGRAM_PTC = 10065

	// The leve l 1_ actstate.

	LEVEL1_ACTSTATE = 10070

	// The leve l 2_ actstate.

	LEVEL1_INPUT_STATE = 10074

	// neu seit 3.06: Akt. Werte Pumpen

	// The currentpum p 1_ actpercent.

	CURRENTPUMP1_ACTPERCENT = 10080

	// The currentpum p 2_ actpercent.

	CURRENTPUMP2_ACTPERCENT = 10081

	// The currentpum p 3_ actpercent.

	CURRENTPUMP3_ACTPERCENT = 10082

	// The currentpum p 4_ actpercent.

	CURRENTPUMP4_ACTPERCENT = 10083

	// Alarm

	// The isalarm.

	ISALARM = 10090

	// The digitalinputsstate.

	DIGITALINPUTSSTATE = 10091

	// Diagnose-/Inbetriebnahmeparameter

	// The lightscenetesttime.

	LIGHTSCENETESTTIME = 10095

	// The r w_ eeprom.

	RW_EEPROM = 10096

	// The opmode.

	OPMODE = 10097

	// The s p 1_ state.

	SP1_STATE = 10100

	// ...

	// The s p 24_ state.

	SP24_STATE = 10123

	// The s p_ al l_ state.

	SP_ALL_STATE = 10126

	SP_ALL_CURRENT = 10127

	// The slotcount.

	SLOTCOUNT = 10129

	// The ke y_ state.

	KEY_STATE = 10130

	// The buzzerstate.

	BUZZERSTATE = 10131

	// The alarmledstate.

	ALARMLEDSTATE = 10132

	// The dcfstate.

	DCFSTATE = 10134

	// The modu l 0 state.

	MODUL0STATE = 10137

	// The modu l 1 state.

	MODUL1STATE = 10138

	// The modu l 2 state.

	MODUL2STATE = 10139

	// The rtctickcount.

	RTCTICKCOUNT = 10140

	// The setlantronixconfi g_ slo t 0.

	SETLANTRONIXCONFIG_SLOT0 = 10141

	// The setlantronixconfi g_ slo t 1.

	SETLANTRONIXCONFIG_SLOT1 = 10142

	// The setlantronixconfi g_ slo t 2.

	SETLANTRONIXCONFIG_SLOT2 = 10143

	// The setlantronixconfigsecurit y_ slo t 0.

	SETLANTRONIXCONFIGSECURITY_SLOT0 = 10144

	// The setlantronixconfigsecurit y_ slo t 1.

	SETLANTRONIXCONFIGSECURITY_SLOT1 = 10145

	// The setlantronixconfigsecurit y_ slo t 2.

	SETLANTRONIXCONFIGSECURITY_SLOT2 = 10146

	// The sendkey.

	SENDKEY = 10150

	// The getdisplaylin e 1.

	GETDISPLAYLINE1 = 10151

	// The getdisplaylin e 2.

	GETDISPLAYLINE2 = 10152

	// The getdisplaystate.

	GETDISPLAYSTATE = 10153

	// The modu l 0 version.

	MODUL0VERSION = 10154

	// The modu l 1 version.

	MODUL1VERSION = 10155

	// The modu l 2 version.

	MODUL2VERSION = 10156

	// The modu l 0 id.

	MODUL0ID = 10157

	// The modu l 1 id.

	MODUL1ID = 10158

	// The modu l 2 id.

	MODUL2ID = 10159

	// Messwerterfassung

	// The measuremen t_ newdatacount.

	MEASUREMENT_NEWDATACOUNT = 10160

	// The measuremen t_ usedmemorysize.

	MEASUREMENT_USEDMEMORYSIZE = 10162

	// The measuremen t_ lastsampletime.

	MEASUREMENT_LASTSAMPLETIME = 10163

	// The measuremen t_ lastsampledate.

	MEASUREMENT_LASTSAMPLEDATE = 10164

	// The measuremen t_ validdatacount.

	MEASUREMENT_VALIDDATACOUNT = 10165

	// The measuremen t_ getdatarecord.

	MEASUREMENT_GETDATARECORD = 10166

	// The measuremen t_ increaddatapointer.

	MEASUREMENT_INCREADDATAPOINTER = 10167

	// The l 1 t o 10 vin t 1_ pwmvalue.

	L1TO10VINT1_PWMVALUE = 10170

	// ...

	// The l 1 t o 10 vin t 10_ pwmvalue.

	L1TO10VINT10_PWMVALUE = 10179

	// Debugging

	// The tes t 1.

	TEST1 = 10200

	// The tes t 2.

	TEST2 = 10201

	// The tes t 3.

	TEST3 = 10202

	// The tes t 4.

	TEST4 = 10203

	// The tes t 5.

	TEST5 = 10204

	// The tes t 6.

	TEST6 = 10205

	// The tes t 7.

	TEST7 = 10206

	// The tes t 8.

	TEST8 = 10207

	// The tes t 9.

	TEST9 = 10208

	// The tes t 10.

	TEST10 = 10209

	// The freestack.

	FREESTACK = 10210

	// View

	// The vie w_ pwmcontrast.

	VIEW_PWMCONTRAST = 10220

	// The sensorpar a 1_ completestring.

	SENSORPARA1_COMPLETESTRING = 10250

	MAINTENANCE_REMATINGTIME = 10270

	MAINTENANCE_ISACTIVE = 10271

	// The sensorpar a 32_ completestring.

	SENSORPARA32_COMPLETESTRING = 10281

	// The aquailluminatio n_ setbasetemp.

	AQUAILLUMINATION_SETBASETEMP = 10320

	// The aquailluminatio n_ setflashint.

	AQUAILLUMINATION_SETFLASHINT = 10321

	// The dal i_ sendcommand.

	DALI_SENDCOMMAND = 10340

	// The dal i_ getlastresponse.

	DALI_GETLASTRESPONSE = 10341

	///#define BLOCKSIZE_ILLUMINATIONSTATES    4

	// The illuminatio n 1_ actpercent.

	ILLUMINATION1_ACTPERCENT = 10350

	// The illuminatio n 1_ ohm.

	ILLUMINATION1_OHM = 10351

	// The illuminatio n 1_ re s 1.

	ILLUMINATION1_RES1 = 10352

	// The illuminatio n 1_ re s 2.

	ILLUMINATION1_RES2 = 10353

	// ...

	// The illuminatio n 8_ re s 2.

	ILLUMINATION8_RES2 = 10381

	// The thunderstor m_ manuflash.

	THUNDERSTORM_MANUFLASH = 10400

	// The progra m_ le d_ light.

	PROGRAM_LED_LIGHT = 10401

	// The moo n_ actphase.

	MOON_ACTPHASE = 10402

	// The setdigpbnumbering.

	SETDIGPBNUMBERING = 10403

	// The setdigpbinistate.

	SETDIGPBINISTATE = 10404

	// The invokespecialfunction.

	INVOKESPECIALFUNCTION = 10406

	// zweiter Code-Block - nur für ProfiLux 3

	// The pabserialnumber.

	PABSERIALNUMBER = 10410

	// The pabvalidationkey.

	PABVALIDATIONKEY = 10411

	// The swmpabdirec t_ i d_ dc.

	SWMPABDIRECT_ID_DC = 10412

	// The swmpabdirec t_ ra w 0.

	SWMPABDIRECT_RAW0 = 10413

	// The swmpabdirec t_ ra w 1.

	SWMPABDIRECT_RAW1 = 10414

	// The pa b_ reset.

	PAB_RESET = 10415

	// The cantransparentmode.

	CANTRANSPARENTMODE = 10416

	// The framtest.

	FRAMTEST = 10417

	// The nictest.

	NICTEST = 10418

	DISPLAY1 = 10460
	DISPLAY2 = 10461
	DISPLAY3 = 10462
	DISPLAY4 = 10463
	DISPLAY5 = 10464
	DISPLAY6 = 10465
	DISPLAY7 = 10465

	// Determine count of available resources

	// The getsensorcount.

	GETSENSORCOUNT = 10500

	// The getswitchcount.

	GETSWITCHCOUNT = 10501

	// The getontetotenvintcount.

	GETONTETOTENVINTCOUNT = 10502

	// The getlevelsensorcount.

	GETLEVELSENSORCOUNT = 10503

	// The getserialinterfacecount.

	GETSERIALINTERFACECOUNT = 10504

	// The getdigitalinputcount.

	GETDIGITALINPUTCOUNT = 10505

	// The getreserve d 1 count.

	GETRESERVED1COUNT = 10506

	// The getreserve d 2 count.

	GETRESERVED2COUNT = 10507

	// The getreserve d 3 count.

	GETRESERVED3COUNT = 10508

	// The getreserve d 4 count.

	GETRESERVED4COUNT = 10509

	// The getilluminationcount.

	GETILLUMINATIONCOUNT = 10510

	// The getremindercount.

	GETREMINDERCOUNT = 10511

	// The gettimercount.

	GETTIMERCOUNT = 10512

	// The getproglogiccount.

	GETPROGLOGICCOUNT = 10513

	// The getcurrentpumpcount.

	GETCURRENTPUMPCOUNT = 10514

	// The getvariableilluminationcount.

	GETVARIABLEILLUMINATIONCOUNT = 10515

	// The co m_ portcount.

	COM_PORTCOUNT = 10600

	SENSOR1_NAME = 18000

	ILLUMINATION1_NAME = 18032

	SWITCHPLUG1_NAME = 18064

	LEVEL_NAME = 18128

	MAINT_NAME = 18144

	// The multiplecodeinf o_0.

	MULTIPLECODEINFO_0 = 20000

	// The multiplecodeinf o_19999.

	MULTIPLECODEINFO_19999 = 39999

// ReSharper restore InconsistentNaming
// ReSharper restore UnusedMember.Global
)
