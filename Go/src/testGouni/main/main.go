package main

import (
	"bytes"
	"net/http"
	"strconv"
	"fmt"
)

const contentType = "Content-Type"
const url = "http://da027-ldng-1qns990.sii24.pole-emploi.intra:19901/da027-diffusionspecifiqueoffre/v2/integrationoffre/poleemploi/"
//const url = "http://da027-ldng-i8rsbc1-1.sii24.pole-emploi.intra:20283/da027-diffusionspecifiqueoffre/v2/integrationoffre/poleemploi/"

func main() {

	// postGouni("T0000")s
	for i := 0; i < 10000; i ++ {
		offre := "A" + strconv.Itoa(i)

		fmt.Println(offre + " : status de la requête => " + postGouni(offre))
	}

}

func postGouni(numOffre string) string {
	var bodyXml string = `<FluxGOUNI>
	<casUtilisation>TP_ERO_CR</casUtilisation>
	<horodatageFlux>20/03/2018 11:51:03.342</horodatageFlux>
	<OffreGOUNI>
	<idOffre>` + numOffre + `</idOffre>
	<codeVersionOffre>G</codeVersionOffre>
	<etatOffre>EC</etatOffre>
	<codeTypoOffre>C</codeTypoOffre>
	<horodatageCreation>20/03/2018 10:50:58.655</horodatageCreation>
	<horodatageModification>20/03/2018 10:50:58.655</horodatageModification>
	<topDiffusionInternet>O</topDiffusionInternet>
	<topDiffusionEpartenet>O</topDiffusionEpartenet>
	<etatPlafondTCD>0</etatPlafondTCD>
	<dateCreationReport>20/03/2018</dateCreationReport>
	<dateModificationReport>20/03/2018</dateModificationReport>
	<unitePriseOffre>62566</unitePriseOffre>
	<dateEffetOffre>20/03/2018</dateEffetOffre>
	<horodatageDernierOSU>20/03/2018 10:50:59.101</horodatageDernierOSU>
	<EmployeurGOUNI>
	<numEtab>28927456</numEtab>
	<urlEntreprise>http://www.lidl.fr` + numOffre + `</urlEntreprise>
	</EmployeurGOUNI>
	<EmploiGOUNI>
	<codeRome>I1203</codeRome>
	<codeAppellation>38086</codeAppellation>
	<ListeActivites>
	<ActiviteGOUNI>
	<codeOGR>106661</codeOGR>
	<competenceExigee>N</competenceExigee>
	</ActiviteGOUNI>
	<ActiviteGOUNI>
	<codeOGR>121493</codeOGR>
	<competenceExigee>N</competenceExigee>
	</ActiviteGOUNI>
	</ListeActivites>
	<ListeCompetences>
	<CompetenceGOUNI>
	<codeOGR>120676</codeOGR>
	<competenceExigee>N</competenceExigee>
	</CompetenceGOUNI>
	<CompetenceGOUNI>
	<codeOGR>100017</codeOGR>
	<competenceExigee>N</competenceExigee>
	</CompetenceGOUNI>
	</ListeCompetences>
	<ListeEnvironnementsTravail/>
	<nbPostesCrees>1</nbPostesCrees>
	<descriptionOffre>desc qksjdlkqsjdkjqskdjqskjdqskjdk pour flux gouni eures tests pour us591
	</descriptionOffre>
	<accesTH>O</accesTH>
	<codeONA>RVR</codeONA>
	<descriptionEntreprise>ljlqsjdkqsjdkljqskldjqskljdklqsjdlkqsjdlkqjsldjqskld pour us 591</descriptionEntreprise>
	<intituleOffre>Service à domicile H/F` + numOffre + `</intituleOffre>
	<ListeCompetencesLibres>
	<CompetenceLibreGOUNI>
	<libelleCompetence>agile</libelleCompetence>
	<competenceExigee>O</competenceExigee>
	</CompetenceLibreGOUNI>
	</ListeCompetencesLibres>
	<ListeQualitesProfessionnelles>
	<QualiteProfessionnelleGOUNI>
	<codeQualiteProfessionnelle>11</codeQualiteProfessionnelle>
	</QualiteProfessionnelleGOUNI>
	<QualiteProfessionnelleGOUNI>
	<codeQualiteProfessionnelle>4</codeQualiteProfessionnelle>
	</QualiteProfessionnelleGOUNI>
	</ListeQualitesProfessionnelles>
	</EmploiGOUNI>
	<PosteGOUNI>
	<codeNatureContrat>E1</codeNatureContrat>
	<codeTypeContrat>CDD</codeTypeContrat>
	<ListeLieuxTravail>
	<LieuTravailGOUNI>
	<typeLieuTravail>CO</typeLieuTravail>
	<libelleLieuTravail>AMBRONAY</libelleLieuTravail>
	<codePaysLieuTravail></codePaysLieuTravail>
	<codeRegionLieuTravail>93</codeRegionLieuTravail>
	<codeDepartementLieuTravail>01</codeDepartementLieuTravail>
	<codeCommuneLieuTravail></codeCommuneLieuTravail>
	<codePostalLieutravail></codePostalLieutravail>
	<codeContinent>00</codeContinent>
	</LieuTravailGOUNI>
	</ListeLieuxTravail>
	<uniteDureeContrat>MO</uniteDureeContrat>
	<dureeContrat>24</dureeContrat>
	<dateDebutContrat>15/04/2018</dateDebutContrat>
	<codeQualification>3</codeQualification>
	<codeConditionExercice>AUT</codeConditionExercice>
	<codeTypeHoraire>NOR</codeTypeHoraire>
	<complTypeHoraire>qsdqsjkdlkqsjlkdjqs</complTypeHoraire>
	<nbHeuresHebdo>38</nbHeuresHebdo>
	<nbMinutesHebdo>0</nbMinutesHebdo>
	<codeFreqDeplacement>3</codeFreqDeplacement>
	<codeTypeDeplacement>1</codeTypeDeplacement>
	<typeSalaire>M</typeSalaire>
	<salaireMin>2000</salaireMin>
	<salaireMax>2500</salaireMax>
	<sigleMonetaire>EU</sigleMonetaire>
	<nbMoisSalaire>12</nbMoisSalaire>
	<salaireAnnuelMini>24000,00</salaireAnnuelMini>
	<salaireAnnuelMaxi>30000,00</salaireAnnuelMaxi>
	<ListeComplementsSalaire>
	<ComplementSalaireGOUNI>
	<codeComplementSalaire>2</codeComplementSalaire>
	</ComplementSalaireGOUNI>
	</ListeComplementsSalaire>
	<topDesQuePossible>N</topDesQuePossible>
	</PosteGOUNI>
	<ProfilGOUNI>
	<codeTypeExpPro>D</codeTypeExpPro>
	<comExpPro>qjsdklqjlkdsjdklqsjd</comExpPro>
	<ListeFormations>
	<FormationGOUNI>
	<codeTypeFormation>NV1</codeTypeFormation>
	<codeDomaineFormation>11001</codeDomaineFormation>
	<intituleDiplomeFormation>sldklmqsdkmlqskld</intituleDiplomeFormation>
	<codeExigibiliteFormation>E</codeExigibiliteFormation>
	<numOrdreFormation>1</numOrdreFormation>
	</FormationGOUNI>
	</ListeFormations>
	<ListeLangues>
	<LangueGOUNI>
	<codeLangue>AL</codeLangue>
	<codeNiveauLangue>1</codeNiveauLangue>
	<codeExigibiliteLangue>E</codeExigibiliteLangue>
	<numOrdreLangue>1</numOrdreLangue>
	</LangueGOUNI>
	</ListeLangues>
	<ListeOutilBureautique/>
	<ListePermis>
	<PermisGOUNI>
	<codePermis>A</codePermis>
	<codeExigibilitePermis>E</codeExigibilitePermis>
	<numOrdrePermis>1</numOrdrePermis>
	</PermisGOUNI>
	</ListePermis>
	</ProfilGOUNI>
	<SuiviOffreGOUNI>
	<codePreselection>C</codePreselection>
	<quotaMERActuel>5</quotaMERActuel>
	<quotaMERInitial>5</quotaMERInitial>
	<codeMPE>TEL</codeMPE>
	<comMPE>jkljqslkfjsdkl</comMPE>
	<comMPA>qsmlkdmqskdlmqs</comMPA>
	<enseignementEtab></enseignementEtab>
	<codeTrancheEffectifEtab>01</codeTrancheEffectifEtab>
	<apeEtab>5630Z</apeEtab>
	<complAdress1Etab>12</complAdress1Etab>
	<complAdress2Etab>A</complAdress2Etab>
	<numVoieEtab>49</numVoieEtab>
	<typeVoieEtab>RUE</typeVoieEtab>
	<nomVoieEtab>NATIONALE</nomVoieEtab>
	<distributionSpeEtab>BP - 15</distributionSpeEtab>
	<acheminementPostalEtab>BOULOGNE SUR MER</acheminementPostalEtab>
	<codePostalEtab>62200</codePostalEtab>
	<codeCommuneEtab>62160</codeCommuneEtab>
	<codeStructAccueil>62566</codeStructAccueil>
	<nomStructAccueil>BOULOGNE-SUR-MER</nomStructAccueil>
	<libelleVoirieStructAccueil>135 BOULEVARD DAUNOU                  </libelleVoirieStructAccueil>
	<codePostalStructAccueil>62200</codePostalStructAccueil>
	<codeCommuneStructAccueil>62160</codeCommuneStructAccueil>
	<telStructAccueil>3949</telStructAccueil>
	<faxStructAccueil>0321100650</faxStructAccueil>
	<mailStructAccueil>ape.62566@pole-emploi.fr</mailStructAccueil>
	<diffusionServeur>I</diffusionServeur>
	<codeEtatDiffusionEuropa>O</codeEtatDiffusionEuropa>
	<codeRegionDiffusionEuropa>FR71</codeRegionDiffusionEuropa>
	<ListeSuivis>
	<SuiviGOUNI>
	<horodatageSuivi>20/03/2018 10:50:59.101</horodatageSuivi>
	<codeCiviliteCorrespALE>1</codeCiviliteCorrespALE>
	<nomCorrespALE>Service Entreprise</nomCorrespALE>
	<prenomCorrespALE>test prenom</prenomCorrespALE>
	<codeCorrespEquipe>1</codeCorrespEquipe>
	<codeAgentCorrespALE>STR</codeAgentCorrespALE>
	<telCorrespALE>0321100641</telCorrespALE>
	<mailCorrespALE>entreprise.npc0144@pole-emploi.net</mailCorrespALE>
	<dateProchainSuivi>27/02/2018</dateProchainSuivi>
	<uniteSuivi>62566</uniteSuivi>
	</SuiviGOUNI>
	</ListeSuivis>
	<ListeCorrespondants>
	<CorrespondantGOUNI>
	<civiliteCorrespondant>1</civiliteCorrespondant>
	<nomCorrespondant>tre</nomCorrespondant>
	<prenomCorrespondant>tre</prenomCorrespondant>
	<telCorrespondant>020202020202022</telCorrespondant>
	<faxCorrespondant>023011353454532</faxCorrespondant>
	<mailCorrespondant>toto.tata@free.fr</mailCorrespondant>
	<numOrdreCorrespondant>1</numOrdreCorrespondant>
	<codeTypeCorrespondant>R</codeTypeCorrespondant>
	</CorrespondantGOUNI>
	</ListeCorrespondants>
	<ListeEditions>
	<EditionGOUNI>
	<horodatageEdition>20/03/2018 10:50:59.102</horodatageEdition>
	<formatEdition>M</formatEdition>
	<dateEdition>20/03/2018</dateEdition>
	<codeCourrierEdite>CC</codeCourrierEdite>
	</EditionGOUNI>
	</ListeEditions>
	<ListeDelegationUniteMER>
	<DelegationUniteMERGOUNI>
	<codeUniteDelegMER>00500</codeUniteDelegMER>
	</DelegationUniteMERGOUNI>
	<DelegationUniteMERGOUNI>
	<codeUniteDelegMER>00503</codeUniteDelegMER>
	</DelegationUniteMERGOUNI>
	</ListeDelegationUniteMER>
	<ListeDelegationBassinMER>
	<DelegationBassinMERGOUNI>
	<codeBassinDelegMER>AL67H</codeBassinDelegMER>
	</DelegationBassinMERGOUNI>
	<DelegationBassinMERGOUNI>
	<codeBassinDelegMER>AL67N</codeBassinDelegMER>
	</DelegationBassinMERGOUNI>
	</ListeDelegationBassinMER>
	<SuiviGOUNI>
	<horodatageSuivi>20/03/2018 10:50:59.101</horodatageSuivi>
	<codeCiviliteCorrespALE>1</codeCiviliteCorrespALE>
	<nomCorrespALE>Service Entreprise</nomCorrespALE>
	<prenomCorrespALE>test prenom</prenomCorrespALE>
	<codeCorrespEquipe>1</codeCorrespEquipe>
	<codeAgentCorrespALE>STR</codeAgentCorrespALE>
	<telCorrespALE>0321100641</telCorrespALE>
	<mailCorrespALE>entreprise.npc0144@pole-emploi.net</mailCorrespALE>
	<dateProchainSuivi>27/02/2018</dateProchainSuivi>
	<uniteSuivi>62566</uniteSuivi>
	</SuiviGOUNI>
	<ServiceRecrutGOUNI>
	<horodatageServiceRecrut>20/03/2018 10:50:59.100</horodatageServiceRecrut>
	<typeServiceRecrut>ACC</typeServiceRecrut>
	<ListeServiceNiveau1>
	<ServiceNiveau1GOUNI>
	<codeServiceNiveau1>PRE</codeServiceNiveau1>
	<ListeServiceNiveau2>
	<ServiceNiveau2GOUNI>
	<codeServiceNiveau2>PRE</codeServiceNiveau2>
	<ListeServiceNiveau3/>
	</ServiceNiveau2GOUNI>
	</ListeServiceNiveau2>
	</ServiceNiveau1GOUNI>
	</ListeServiceNiveau1>
	</ServiceRecrutGOUNI>
	<ListeServiceRecrut>
	<ServiceRecrutGOUNI>
	<horodatageServiceRecrut>20/03/2018 10:50:59.100</horodatageServiceRecrut>
	<typeServiceRecrut>ACC</typeServiceRecrut>
	<ListeServiceNiveau1>
	<ServiceNiveau1GOUNI>
	<codeServiceNiveau1>PRE</codeServiceNiveau1>
	<ListeServiceNiveau2>
	<ServiceNiveau2GOUNI>
	<codeServiceNiveau2>PRE</codeServiceNiveau2>
	<ListeServiceNiveau3/>
	</ServiceNiveau2GOUNI>
	</ListeServiceNiveau2>
	</ServiceNiveau1GOUNI>
	</ListeServiceNiveau1>
	</ServiceRecrutGOUNI>
	</ListeServiceRecrut>
	<codeDistributionSpeEtab>BP</codeDistributionSpeEtab>
	<libDistributionSpeEtab>15</libDistributionSpeEtab>
	<codePaysEtab>01</codePaysEtab>
	</SuiviOffreGOUNI>
	<ListeStatistiques>
	<StatistiqueGOUNI>
	<nbPostesRestant>1</nbPostesRestant>
	<nbPostesEnAjoutDepuisCreation>0</nbPostesEnAjoutDepuisCreation>
	<nbPostesSatisfaitsHorsMER>0</nbPostesSatisfaitsHorsMER>
	<nbPostesSatisfaitsParMER>0</nbPostesSatisfaitsParMER>
	<nbPostesPourvusEnInterne>0</nbPostesPourvusEnInterne>
	<nbPostesDisparus>0</nbPostesDisparus>
	<nbPostesAbandonnes>0</nbPostesAbandonnes>
	<nbMERTotal>0</nbMERTotal>
	<nbTotalSuiviOffre>0</nbTotalSuiviOffre>
	<topRepriseOffre>0</topRepriseOffre>
	<topOffresCreeesParDOL>0</topOffresCreeesParDOL>
	<topOffresSuspenduesParDOL>0</topOffresSuspenduesParDOL>
	<dateDernierSuiviOffre>20/03/2018</dateDernierSuiviOffre>
	<nbMECEnCoursOblig>0</nbMECEnCoursOblig>
	</StatistiqueGOUNI>
	</ListeStatistiques>
	<ListeMajs>
	<MAJGOUNI>
	<horodatageMAJ>20/03/2018 10:50:59.097</horodatageMAJ>
	<typeMAJ>PDI</typeMAJ>
	<dateMAJ>20/03/2018</dateMAJ>
	<codeTypeActeurMAJ>CO</codeTypeActeurMAJ>
	<serviceOrigineOffre>HU</serviceOrigineOffre>
	</MAJGOUNI>
	<MAJGOUNI>
	<horodatageMAJ>20/03/2018 10:50:59.098</horodatageMAJ>
	<typeMAJ>CRE</typeMAJ>
	<dateMAJ>20/03/2018</dateMAJ>
	<codeTypeActeurMAJ>CO</codeTypeActeurMAJ>
	<serviceOrigineOffre>HU</serviceOrigineOffre>
	</MAJGOUNI>
	<MAJGOUNI>
	<horodatageMAJ>20/03/2018 10:50:59.099</horodatageMAJ>
	<typeMAJ>MAQ</typeMAJ>
	<dateMAJ>20/03/2018</dateMAJ>
	<codeTypeActeurMAJ>CO</codeTypeActeurMAJ>
	<serviceOrigineOffre>HU</serviceOrigineOffre>
	<codeMotifMAJ>QOK</codeMotifMAJ>
	</MAJGOUNI>
	<MAJGOUNI>
	<horodatageMAJ>20/03/2018 10:51:00.662</horodatageMAJ>
	<typeMAJ>MOC</typeMAJ>
	<dateMAJ>20/03/2018</dateMAJ>
	<codeTypeActeurMAJ>CO</codeTypeActeurMAJ>
	<serviceOrigineOffre>HU</serviceOrigineOffre>
	<codeMotifMAJ>COK</codeMotifMAJ>
	</MAJGOUNI>
	</ListeMajs>
	<typeAffichage>N</typeAffichage>
	<horodatageDernierOSR>20/03/2018 10:50:59.100</horodatageDernierOSR>
	<numVersion>1</numVersion>
	<LibellesDmcGOUNI>
	<libAppellation>Agent / Agente d'entretien travaux à domicile</libAppellation>
	<libROME>Maintenance des bâtiments et des locaux</libROME>
	<libelleLieuTravail>01 - AMBRONAY</libelleLieuTravail>
	<libContrat>Contrat à durée déterminée - 24 Mois</libContrat>
	<libNatureContrat>Contrat travail</libNatureContrat>
	<libExperience>Débutant accepté - qjsdklqjlkdsjdklqsjd</libExperience>
	<libFormation1>Bac+5 et plus ou équivalent</libFormation1>
	<libDomaineFormation1>Théorie graphes exigé</libDomaineFormation1>
	<libLangue1>Allemand Courant Exigé</libLangue1>
	<libPermis1>A - Moto Exigé</libPermis1>
	<libQualification>Ouvrier qualifié (P1,P2)</libQualification>
	<libSalaire>Mensuel de 2000 Euros à 2500 Euros sur 12 mois</libSalaire>
	<libComplementSalaire1>Véhicule</libComplementSalaire1>
	<libDureeTravail>38H Horaires normaux</libDureeTravail>
	<libDeplacement>Fréquents Départemental</libDeplacement>
	<libApeEtab>Débits de boissons</libApeEtab>
	<libTrancheEffectifEtab>0 salarié</libTrancheEffectifEtab>
	<libEnseigneOffreur>LE RALLYE
	ljlqsjdkqsjdkljqskldjqskljdklqsjdlkqsjdlkqjsld
	jqskld
	M. tre tre
	toto.tata@free.fr
	jkljqslkfjsdkl</libEnseigneOffreur>
	<libContactPropOffreur>LE RALLYE - M. tre tre</libContactPropOffreur>
	<contactPropLigne1>Courriel : toto.tata@free.fr</contactPropLigne1>
	<contactPropCommentaire>jkljqslkfjsdkl</contactPropCommentaire>
	</LibellesDmcGOUNI>
	</OffreGOUNI>
	</FluxGOUNI>`

	body := bytes.NewBufferString(bodyXml)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("pe-id-utilisateur","bench")
	req.Header.Set("pe-nom-application","eures")
	req.Header.Set("pe-nom-domaine","DA027")
	req.Header.Set("pe-id-environnement","ENV_IQRFR_DA027")
	req.Header.Set("pe-id-correlation","1-98796987")
	req.Header.Set("Accept","application/xml")
	req.Header.Set("Content-Type","application/xml")

	if nil != err {
		panic(err)
	}

	res, _ := client.Do(req)

	return res.Status
}
