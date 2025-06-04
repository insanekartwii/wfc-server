package gpcm

var (
	WWFCMsgUnknownLoginError = WWFCErrorMessage{
		ErrorCode: 22000,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"IKWFCへの ログイン中に\n" +
				"不明なエラー が発生しました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"An unknown error has occurred\n" +
				"while logging in to IKWFC.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Ein unbekannter Fehler ist beim\n" +
				"Verbinden mit IKWFC aufgetreten.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Un error desconocido ha ocurrido\n" +
				"al conectarse a IKWFC.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"È stato riscontrato un errore sconosciuto\n" +
				"durante l'accesso alla IKWFC.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Er is een onbekende fout opgetreden\n" +
				"tijdens het verbinden met IKWFC.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"IKWFC에 연결 도중\n" +
				"알 수 없는 오류가 발생했습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Při přihlašování k IKWFC\n" +
				"došlo k neznámé chybě.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Во время входа в IKWFC\n" +
				"произошла ошибка.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Retro Rewind'a giriş yaparken\n" +
				"bilinmeyen bir hata oluştu.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Une erreur inconnue s'est produite\n" +
				"pendant la connexion à IKWFC.\n" +
				"\n" +
				"Code Erreur:  %[1]d",
		},
	}

	WWFCMsgDolphinSetupRequired = WWFCErrorMessage{
		ErrorCode: 22001,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Dolphinで IKWFCに接続するには\n" +
				"ついかの セットアップが必要です\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"Additional setup is required\n" +
				"to use IKWFC on Dolphin.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Für die Verwendung von IKWFC auf Dolphin\n" +
				"ist eine zusätzliche Einrichtung erforderlich.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Se requiere una instalación adicional\n" +
				"para poder usar IKWFC en Dolphin.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"Un'ulteriore installazione è necessaria\n" +
				"per usare la IKWFC su Dolphin.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Verdere configuraties zijn vereist\n" +
				"om IKWFC op Dolphin te gebruiken.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"돌핀에서 IKWFC를 사용하려면\n" +
				"추가 작업이 필요합니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"K použití IKWFC na Dolphinu\n" +
				"je nutné další nastavení.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Чтобы играть по IKWFC\n" +
				"в Dolphin, необходимо провести\n" +
				"дополнительную настройку.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"IKWFC'yi Dolphin'de kullanabilmek\n" +
				"için birkaç adım daha gerekiyor.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Une installation additionnelle est requise\n" +
				"pour utiliser IKWFC sur Dolphin.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgProfileBannedTOS = WWFCErrorMessage{
		ErrorCode: 22002,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"利用きやくに いはんしたため\n" +
				"IKWFCから BANされました\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from IKWFC\n" +
				"due to a violation of the\n" +
				"Terms of Service.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von IKWFC\n" +
				"wegen eines Verstoßes der\n" +
				"Terms of Service gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Te han baneado de IKWFC\n" +
				"debido a una violación de los\n" +
				"Terminos de Servicio.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalla IKWFC\n" +
				"a causa di una violazione dei\n" +
				"Termini di Servizio.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van IKWFC\n" +
				"vanwege een overtreding van de\n" +
				"gebruiksvoorwaarden.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangKorean: "" +
				"이용약관 위반으로\n" +
				"IKWFC 계정이\n" +
				"정지됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz IKWFC\n" +
				"z důvodu porušení\n" +
				"Podmínek Služby.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вам запрещено играть в Retro\n" +
				"WFC из-за нарушения условий\n" +
				"использования сервиса.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Hizmet Şartlarını ihlal ettiğinizden\n" +
				"dolayı IKWFC'ye erişiminiz\n" +
				"yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) de IKWFC\n" +
				"à cause d'une violation des\n" +
				"Conditions de Service.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
		},
	}

	WWFCMsgProfileBannedTOSNow = WWFCErrorMessage{
		ErrorCode: 22002,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"利用きやくに いはんしたため\n" +
				"IKWFCから BANされています\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from IKWFC\n" +
				"due to a violation of the\n" +
				"Terms of Service.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von IKWFC\n" +
				"wegen eines Verstoßes der\n" +
				"Terms of Service gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Te han baneado de IKWFC\n" +
				"debido a una violación de los\n" +
				"Terminos de Servicio.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalla\n" +
				"IKWFC a causa di una violazione\n" +
				"dei Termini di Servizio.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van IKWFC\n" +
				"vanwege een overtreding van de\n" +
				"gebruiksvoorwaarden.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangKorean: "" +
				"이용약관 위반으로\n" +
				"IKWFC 계정이\n" +
				"정지됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz IKWFC\n" +
				"z důvodu porušení\n" +
				"Podmínek Služby.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Отныне вам запрещено играть\n" +
				"в IKWFC из-за нарушения\n" +
				"условий использования сервиса.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Hizmet Şartlarını ihlal ettiğinizden\n" +
				"dolayı IKWFC'ye erişiminiz\n" +
				"yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) de IKWFC\n" +
				"à cause d'une violation des\n" +
				"Conditions de Service.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestricted = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"IKWFCの ルールにいはんしたため\n" +
				"オンライン対戦から BANされました\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from public\n" +
				"matches due to a violation\n" +
				"of the IKWFC Rules.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentl. Räumen\n" +
				"wegen eines Verstoßes der\n" +
				"IKWFC Regeln gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Te han baneado de partidas públicas\n" +
				"debido a una violación de las\n" +
				"reglas de IKWFC.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalle corse\n" +
				"pubbliche a causa di una violazione\n" +
				"delle regole della IKWFC.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare\n" +
				"wedstrijden vanwege een overtreding\n" +
				"van de IKWFC-regels.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangKorean: "" +
				"IKWFC 규정 위반으로\n" +
				"공개 경기에서 차단됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných\n" +
				"zápasů z důvodu porušení\n" +
				"pravidel IKWFC.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вам запрещено участвовать\n" +
				"в публичных играх из-за\n" +
				"нарушения правил IKWFC.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"IKWFC kurallarını ihlal\n" +
				"ettiğinizden dolayı herkese\n" +
				"açık yarışlara erişiminiz yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) des matchs\n" +
				"public à cause d'un violation d'une\n" +
				"des règles de IKWFC.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestrictedNow = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"IKWFCの ルールにいはんしたため\n" +
				"オンライン対戦から BANされています\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You have been banned from public\n" +
				"matches due to a violation\n" +
				"of the IKWFC Rules.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentl. Räumen\n" +
				"wegen eines Verstoßes der\n" +
				"IKWFC Regeln gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x,",
			LangSpanish: "" +
				"Te han baneado de partidas públicas\n" +
				"debido a una violación de las\n" +
				"reglas de IKWFC.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalle corse\n" +
				"pubbliche a causa di una violazione\n" +
				"delle regole della IKWFC.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare\n" +
				"wedstrijden vanwege een overtreding\n" +
				"van de IKWFC-regels.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangKorean: "" +
				"IKWFC 규정 위반으로\n" +
				"공개 경기에서 차단됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných\n" +
				"zápasů z důvodu porušení\n" +
				"pravidel IKWFC.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Отныне вам запрещено участвовать\n" +
				"в публичных играх из-за нарушения\n" +
				"правил IKWFC.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"IKWFC kurallarını ihlal\n" +
				"ettiğinizden dolayı herkese\n" +
				"açık yarışlara erişiminiz yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) des matchs\n" +
				"publics à cause d'un violation d'une\n" +
				"des règles de IKWFC.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestrictedCustom = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"オンライン対戦から BANされました\n" +
				"りゆう： %[3]s\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from public matches.\n" +
				"Reason: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentlichen Matches ausgeschlossen.\n" +
				"Grund: %[3]s\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Estás baneado de partidas públicas.\n" +
				"Motivo: %[3]s\n" +
				"Código de Error: %[1]d\n" +
				"Info. de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei bannato dalle corse pubbliche.\n" +
				"Motivo: %[3]s\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare wedstrijden.\n" +
				"Reden: %[3]s\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangKorean: "" +
				"공개 경기에서 차단됐습니다.\n" +
				"사유: %[3]s\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných zápasů.\n" +
				"Důvod: %[3]s\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вам запрещено участвовать\n" +
				"в публичных играх.\n" +
				"\n" +
				"Причина: %[3]s\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Herkese açık yarışlara\n" +
				"erişiminiz yasaklanmıştır.\n" +
				"Sebep: %[3]s\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous êtes banni(e) des matchs publics.\n" +
				"Raison: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Information Support: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestrictedNowCustom = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"オンライン対戦から BANされました\n" +
				"りゆう： %[3]s\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from public matches.\n" +
				"Reason: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentlichen Matches ausgeschlossen.\n" +
				"Grund: %[3]s\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Estás baneado de partidas públicas.\n" +
				"Motivo: %[3]s\n" +
				"Código de Error: %[1]d\n" +
				"Info. de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei bannato dalle corse pubbliche.\n" +
				"Motivo: %[3]s\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare wedstrijden.\n" +
				"Reden: %[3]s\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangKorean: "" +
				"공개 경기에서 차단됐습니다.\n" +
				"사유: %[3]s\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných zápasů.\n" +
				"Důvod: %[3]s\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Отныне вам запрещено\n" +
				"участвовать в публичных играх.\n" +
				"\n" +
				"Причина: %[3]s\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Herkese açık yarışlara\n" +
				"erişiminiz yasaklanmıştır.\n" +
				"Sebep: %[3]s\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous êtes banni(e) des matchs publics.\n" +
				"Raison: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Information Support: NG%08[2]x",
		},
	}

	WWFCMsgKickedGeneric = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"IKWFCから キックされました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from\n" +
				"IKWFC.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest aus IKWFC\n" +
				"gekickt.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Te han expulsado de IKWFC.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso\n" +
				"dalla IKWFC.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit IKWFC\n" +
				"geschopt.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"IKWFC에서 추방됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Byl jsi vyhozen z\n" +
				"IKWFC.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Вас выгнали из IKWFC.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"IKWFC'den atıldınız.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de\n" +
				"IKWFC.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgKickedModerator = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"IKWFCの モデレーターから\n" +
				"キックされました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from\n" +
				"IKWFC by a moderator.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest von einem Moderator\n" +
				"aus IKWFC gekickt.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Un moderador te ha\n" +
				"expulsado de IKWFC.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso dalla\n" +
				"IKWFC da un moderatore.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit IKWFC\n" +
				"geschopt door een moderator.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"관리자에 의해 IKWFC에서\n" +
				"추방됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Byl jsi vyhozen z\n" +
				"IKWFC moderátorem.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Модератор выгнал вас\n" +
				"из IKWFC.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Bir moderatör tarafından\n" +
				"IKWFC'den atıldınız.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de\n" +
				"IKWFC par un modérateur.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgKickedRoomHost = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"フレンドルームの ホストから\n" +
				"キックされました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from the\n" +
				"friend room by the room creator.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest von dem Raum-Ersteller\n" +
				"aus der Freundes-Lobby gekickt.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"El creador de la sala te ha\n" +
				"expulsado de ella.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso dalla\n" +
				"stanza dal suo creatore.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit de vriendenkamer\n" +
				"geschopt door de gastheer.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"방장에 의해 추방됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Hostitel tě vyhodil\n" +
				"z místnosti.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Создатель группы выгнал вас.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Oda kurucusu tarafından\n" +
				"odadan atıldınız.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de la salle\n" +
				"par le créateur.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgKickedCustom = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"IKWFCから キックされました\n" +
				"りゆう： %[3]s\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from\n" +
				"IKWFC.\n" +
				"Reason: %[3]s\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest aus IKWFC\n" +
				"gekickt.\n" +
				"Grund: %[3]s\n" +
				"Fehlercode: %[1]d,",
			LangSpanish: "" +
				"Te han expulsado de IKWFC.\n" +
				"Motivo: %[3]s\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso\n" +
				"dalla IKWFC.\n" +
				"Motivo: %[3]s\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit IKWFC\n" +
				"geschopt.\n" +
				"Reden: %[3]s\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"IKWFC에서 추방됐습니다.\n" +
				"\n" +
				"사유: %[3]s\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Máš zákaz veřejných zápasů.\n" +
				"Důvod: %[3]s\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вас выгнали из IKWFC.\n" +
				"\n" +
				"Причина: %[3]s\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"IKWFC'den atıldınız.\n" +
				"\n" +
				"Sebep: %[3]s\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de\n" +
				"IKWFC.\n" +
				"Raison: %[3]s\n" +
				"Error Code: %[1]d",
		},
	}

	WWFCMsgConsoleMismatch = WWFCErrorMessage{
		ErrorCode: 22005,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"使われているコンソールは\n" +
				"このプロファイルが 登録されたときに\n" +
				"使われたコンソールでは ありません\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The console you are using is not\n" +
				"the device used to register this\n" +
				"profile.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Die Konsole die du gerade nutzt\n" +
				"ist nicht die selbe mit der dieses\n" +
				"Profil erstellt wurde.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"La consola que estas usando no es el\n" +
				"dispositivo usado para registrar este\n" +
				"perfil.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"La console che stai usando non è\n" +
				"il dispositivo usato per\n" +
				"registrare questo profilo.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De console die je gebruikt is niet\n" +
				"het apparaat dat is gebruikt om dit\n" +
				"profiel te registreren.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"사용중인 콘솔이 프로필에\n" +
				"등록된 기기 정보와 다릅니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Konzole, kterou používáš, není\n" +
				"zařízením používaným k registraci\n" +
				"tohoto profilu.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Ваш профиль был зарегистрирован\n" +
				"на другой консоли.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kullandığınız konsol, kaydolduğunuz\n" +
				"profille aynı değil.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"La console que vous utilisé\n" +
				"n'est pas l'appareil utilisé pour\n" +
				"enregistrer ce profil.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgConsoleMismatchDolphin = WWFCErrorMessage{
		ErrorCode: 22005,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"使われているコンソールは\n" +
				"このプロファイルが 登録されたときに\n" +
				"使われたコンソールでは ありません\n" +
				"NANDがただしく 設定されていることを\n" +
				"ご確認ください\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The console you are using is not\n" +
				"the device used to register this\n" +
				"profile. Please make sure you've\n" +
				"set up your NAND correctly.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Die Konsole die du gerade nutzt\n" +
				"ist nicht die selbe mit der dieses\n" +
				"Profil erstellt wurde. Bitte gehe sicher\n" +
				"dass du das NAND korrekt initialisiert hast.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"La consola que estas usando no es el\n" +
				"dispositivo usado para registrar este\n" +
				"perfil. Asegurate que has configurado\n" +
				"correctamente tu NAND.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"La console che stai usando non è\n" +
				"il dispositivo usato per registrare\n" +
				"questo profilo. Assicurati di avere\n" +
				"impostato la tua NAND correttamente.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De console die je gebruikt is niet\n" +
				"het apparaat dat is gebruikt om dit\n" +
				"profiel te registreren. Zorg ervoor dat de\n" +
				"NAND juist is ingesteld.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"사용중인 콘솔이 프로필에\n" +
				"등록된 기기가 아닙니다.\n" +
				"NAND가 제대로 설치 됐는지\n" +
				"확인해 주십시오.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Konzole, kterou používáš, není\n" +
				"zařízením používaným k registraci\n" +
				"tohoto profilu. Ujisti se prosím, že\n" +
				"jsi správně nastavil NAND.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Ваш профиль был зарегистрирован\n" +
				"на другой консоли. Проверьте, что\n" +
				"память NAND настроена правильно.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kullandığınız konsol, kaydolduğunuz\n" +
				"profille aynı değil.\n" +
				"NAND'inizi doğru ayarladığınıza emin olun.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"La consonle que vous utilisé\n" +
				"n'est pas l'appareil utilisé pour\n" +
				"enregistrer ce profil. Assurez-vous d'avoir\n" +
				"configuré votre NAND correctement.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgProfileIDInvalid = WWFCErrorMessage{
		ErrorCode: 22006,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"あなたが登録しようとした\n" +
				"プロファイルIDは むこうです\n" +
				"新しくライセンスをつくりなおしてください\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The profile ID you are trying to\n" +
				"register is invalid.\n" +
				"Please create a new license.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Die Profil-ID die du versuchst zu\n" +
				"registrieren ist ungültig.\n" +
				"Bitte erstelle ein neues Profil.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"El perfil que está tratando de\n" +
				"registrar es invalido.\n" +
				"Cree una nueva licencia.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"L'ID del profilo che stai cercando\n" +
				"di registrare non è valido.\n" +
				"Crea una nuova patente e riprova.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Het profiel-ID dat je probeert te\n" +
				"registreren is ongeldig.\n" +
				"Maak een nieuw profiel aan.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"사용하시려는 프로필 ID는\n" +
				"등록하실 수 없습니다.\n" +
				"새로운 라이센스를 만드십시오.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Profil, který se pokoušíš\n" +
				"zaregistrovat, je neplatný.\n" +
				"Vytvoř si prosím nový účet.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"ID профиля, который вы пытаетесь\n" +
				"зарегистрировать, — некорректный.\n" +
				"Создайте новое удостоверение.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kaydolmak istediğiniz profil ID'si\n" +
				"geçerli bir ID değil.\n" +
				"Lütfen yeni bir lisans oluşturun.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"L'ID du profil que vous essayez\n" +
				"d'enregistrer est invalide.\n" +
				"Veuillez créer une nouveau permis.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgProfileIDInUse = WWFCErrorMessage{
		ErrorCode: 22007,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"あなたが登録しようとした\n" +
				"フレンドコードは すでに登録されています\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The friend code you are trying to\n" +
				"register is already in use.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Den Freundescode den du gerade\n" +
				"registrierst wird bereits verwendet.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"La clave de amigo que está tratando\n" +
				"de registrar, ya está en uso.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"Il codice amico che stai cercando\n" +
				"di registrare è già in uso.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De vriendcode die je probeert te\n" +
				"registreren is al in gebruik.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"해당 친구 코드는 이미 사용중입니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Kód kamaráda, který se pokoušíš\n" +
				"zaregistrovat, se už používá.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Код друга, который вы пытаетесь\n" +
				"зарегистрировать, уже занят.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kaydolmak istediğiniz arkadaş kodu\n" +
				"zaten kullanımda.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Le code ami que vous essayez\n" +
				"d'enregistrer est déjà utilisé.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgPayloadInvalid = WWFCErrorMessage{
		ErrorCode: 22008,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"IKWFCの ペイロードがむこうです\n" +
				"ゲームを 再起動してください\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The IKWFC payload is invalid.\n" +
				"Try restarting your game.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Der IKWFC payload ist ungültig.\n" +
				"Versuche das Spiel neu zu starten.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"IKWFC no cargó correctamente\n" +
				"Intente reiniciar su juego.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"Il payload della IKWFC non è valido.\n" +
				"Prova a riavviare il gioco.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De IKWFC-payload is ongeldig.\n" +
				"Probeer het spel opnieuw op te starten.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"IKWFC 페이로드가 잘못됐습니다.\n" +
				"게임을 재시작 하십시오.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Datový obsah IKWFC je neplatné.\n" +
				"Zkus restartovat hru.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Запущен некорректный пейлоад.\n" +
				"IKWFC. Перезапустите игру.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"IKWFC payloud'u geçerli değil.\n" +
				"Oyunu yeniden başlatmayı deneyin.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Le payload IKWFC est invalide.\n" +
				"Veuillez redémarrer votre jeu.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgInvalidELO = WWFCErrorMessage{
		ErrorCode: 22009,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"VRまたはBRの値が むこうなため\n" +
				"IKWFCから 切断されました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You were disconnected from\n" +
				"IKWFC due to an invalid\n" +
				"VR or BR value.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Deine Verbindung zu IKWFC\n" +
				"durch einen ungültigen VR oder BR\n" +
				"Wert beendet.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Fuiste desconectado debido a discrepancias\n" +
				"con tu valor de PC o PB.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"Sei stato disconnesso dalla IKWFC\n" +
				"a causa di un valore non valido\n" +
				"di punti corsa o punti battaglia.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je verbinding met IKWFC is verbroken\n" +
				"vanwege een ongeldige rp- of gp-waarde.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangKorean: "" +
				"잘못된 VR 또는 BR 값으로 인해\n" +
				"IKWFC 연결이 끊어졌습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Byl jsi odpojen od IKWFC\n" +
				"z důvodu neplatné hodnoty\n" +
				"ZH nebo BH.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Вас отключили от IKWFC\n" +
				"из-за некорректного значения\n" +
				"ГР или БР.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Hatalı bir KP veya SP\n" +
				"değerinden dolayı IKWFC'ye\n" +
				"bağlantınız kesildi.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été déconnecté de\n" +
				"IKWFC à cause d'une valeur invalide\n" +
				"de Points Course ou Points Bataille.\n" +
				"\n" +
				"Code Erreur: %[1]d",
		},
	}

	WWFCMsgInvalidHash = WWFCErrorMessage{
		ErrorCode: 22010,
		MessageRMC: map[byte]string{
			LangEnglish: "" +
				"Invalid pack version!\n" +
				"Please update or reinstall your pack\n" +
				"to log in.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Ungültige Pack-Version!\n" +
				"Bitte update oder installiere\n" +
				"das Pack neu um dich einzuloggen.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangDutch: "" +
				"Ongeldige modpack versie!\n" +
				"Update of herinstalleer de modpack\n" +
				"om in te loggen.\n" +
				"\n" +
				"Foutcode: %[1]d",
		},
	}

)
