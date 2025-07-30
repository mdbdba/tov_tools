package character

import "tov_tools/pkg/static_data"

var Backgrounds = map[string]Background{
	"adherent": {
		Name:             "Adherent",
		BackgroundSource: "Players Guide, pg 118",
		Description: "Before you began adventuring, you committed yourself to a faith, belief, or cause. The " +
			"exacting tasks required of this commitment—daily prayers, holy rites, or cryptic " +
			"ceremonies—instilled in you a sense of duty and purpose. Perhaps you were a hopeful inductee " +
			"into the war god's clergy, a priest excommunicated from a fiend-worshipping sect, or a " +
			"lifelong member of a secret society with global reach. In any case, you still carry the " +
			"teachings and traditions of your devotion.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"History", "Investigation", "Religion", "Persuasion"},
			},
		},
		AdditionalProficiencies: []string{"artist tools"},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"tools": {
				NumberToSelect: 1,
				Options:        static_data.GetToolsetNames(),
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "holy symbol", Quantity: 1},
			{Name: "incense", Quantity: 1},
			{Name: "vestments", Quantity: 1},
			{Name: "clothes common", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 10},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"field medic", "mental fortitude", "ritualist"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "I can test the limits of my devotion out in the wider world through adventuring.",
				2: "Adventuring allows me to learn about and report on other religions and orders.",
				3: "Adventuring frees me to practice more unorthodox methods of worship.",
				4: "I may find others sworn to my order when I am out adventuring.",
				5: "Encountering new people while adventuring lets me share my faith with heretics, pagans, and the uninitiated.",
				6: "When I triumph through adventuring, I will bring glory and notoriety to my order.",
				7: "Adventuring furnishes me with the tithe my order deserves.",
				8: "Staying on the move keeps me from being dragged back to the order from which I narrowly escaped.",
			},
		},
	},
	"artist": {
		Name:             "Artist",
		BackgroundSource: "Players Guide, pg 119",
		Description: "You doggedly practiced artistic pursuits before taking up the adventuring life. " +
			"Countless hours of practice, reflection, and expression altered the way you see the world, " +
			"and demand for your artistic expression grew. Perhaps you began an acrobat honing your body, " +
			"a shadow puppeteer hungry for monstrous material, a dour thespian in search of a patron to " +
			"admire your dark performances, or a bubbly singer enraptured by the applause of strangers. " +
			"No matter what form your expression has taken, you still thrive where an audience waits to " +
			"be entertained, frightened, or inspired.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"Acrobatics", "Insight", "Performance", "Persuasion"},
			},
		},
		AdditionalProficiencies: []string{"artist tools"},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"tools": {
				NumberToSelect: 1,
				Options:        static_data.GetToolsetNames(),
			},
			"languages": {
				NumberToSelect: 1,
				Options:        LanguageNames(),
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "mirror", Quantity: 1},
			{Name: "ink", Quantity: 1},
			{Name: "pen", Quantity: 1},
			{Name: "clothes fine", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 4},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"quick", "scrutinous", "trade skills"},
			},
		},
		Motivations: map[string]map[int]string{
			"artistic": {
				1:  "Painting",
				2:  "Sculpture",
				3:  "Poetry",
				4:  "Storytelling",
				5:  "Acting",
				6:  "Dancing",
				7:  "Juggling",
				8:  "Puppetry",
				9:  "Music",
				10: "Gymnastics",
			},
			"adventuring": {
				1: "Adventuring secures my fortune while my art secures my reputation.",
				2: "Adventuring inspires me by allowing me to meet new people and experience new places.",
				3: "The thrills and terror of adventuring make me far more comfortable in front of even hostile crowds.",
				4: "Adventuring develops skills for me to use when entertaining a crowd.",
				5: "Adventuring puts me out of reach of the patron I've neglected.",
				6: "My mentor was an adventurer. If their travels made them a master, it may work for me as well!",
				7: "Adventuring is how I will find someone who appreciates my art as much as it deserves.",
				8: "Tales of heroism born from adventuring will make my fans hungry for more of my art.",
			},
		},
	},
	"courtier": {
		Name:             "Courtier",
		BackgroundSource: "Players Guide, pg 120",
		Description: "You spent a great deal of time in a royal court. Lessons of decorum and expectations " +
			"of duty and honor granted expert understanding of the world and your place within it, as well " +
			"as the station and manner of others. Perhaps you were a dignitary from a far-off land, a " +
			"master of spies working at the queen's bidding, a constable tasked with capturing fugitives, " +
			"or a noble stricken with wanderlust. Regardless of your courtly appointment, your life was " +
			"one of leadership, service, or privilege, driven by the expectations of others and your own " +
			"ambitions. You still carry these with you.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"History", "Religion", "Insight", "Deception"},
			},
		},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"languages": {
				NumberToSelect: 1,
				Options:        LanguageNames(),
			},
			"tools": {
				NumberToSelect: 1,
				Options:        []string{"artist tools", "navigator tools"},
			},
			"instruments": {
				NumberToSelect: 1,
				Options:        []string{"musical instrument"},
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "writ of nobility", Quantity: 1},
			{Name: "signet ring", Quantity: 1},
			{Name: "clothes fine", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 12},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"combat conditioning", "mental fortitude", "polyglot"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "Adventuring allows me to experience the world firsthand, without court drama.",
				2: "Adventuring is how I will attain glory and stand apart from others of my station.",
				3: "Adventuring is a means to amass power and influence, which I'll use to found my own kingdom.",
				4: "I have been cast out from royal court in disgrace. Adventuring is my best way to find redemption.",
				5: "Danger is my obsession, and adventure is how I'll slake my thirst for it.",
				6: "I have been ordered to adventure by royal decree, and so I shall, until summoned back to court.",
				7: "Through adventuring, I shall forge new alliances for the benefit of the realm.",
				8: "I wish to destroy another member of the court—perhaps adventuring will provide me the means to do so.",
			},
		},
	},
	"criminal": {
		Name:             "Criminal",
		BackgroundSource: "Players Guide, pg 120",
		Description: "You were a cutpurse, grifter, thief, or assassin. Surviving in the criminal underworld " +
			"while plying your nefarious trade taught you patience, resourcefulness, and careful planning. " +
			"Perhaps you were a pickpocket jailed one too many times, a con artist hoping to fleece nobles " +
			"out of their ill-gotten gains, or an assassin ready to turn over a new leaf after being left " +
			"for dead. Regardless, a life of crime has left you tied to society's underbelly.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"Stealth", "Investigation", "Insight", "Deception"},
			},
		},
		AdditionalProficiencies: []string{"thieves cant"},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"tools": {
				NumberToSelect: 1,
				Options:        static_data.GetToolsetNames(),
			},
			"vehicles": {
				NumberToSelect: 1,
				Options:        []string{"land vehicles"},
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "chalk", Quantity: 5},
			{Name: "grappling hook", Quantity: 1},
			{Name: "clothes traveler dark", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 10},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"covert", "scrutinous", "touch of luck"},
			},
		},
		Motivations: map[string]map[int]string{
			"secret": {
				1:  "I inherited a massive fortune but lost it all.",
				2:  "My closest friend wants me dead, and I deserve it.",
				3:  "I am not who I claim to be. I borrowed this life from a dying criminal.",
				4:  "I have a rare terminal illness, and it's making me reckless.",
				5:  "My loving family regularly sends me messages begging me to come home.",
				6:  "I only pretend to be a criminal to make life exciting.",
				7:  "I sometimes make mistakes on purpose just so things will turn violent.",
				8:  "I am terrified of the person I was becoming and pray that I can still change.",
				9:  "I used to be a royal spy until my sovereign betrayed me.",
				10: "I plan to manipulate adventurers who trust me into destroying the enemies of my guild.",
				11: "My family doesn't know I am a criminal, and I'll kill to keep that secret.",
				12: "I am a celebrity in some parts of the world.",
			},
			"adventuring": {
				1: "The bounty on my head is too high! I adventure to keep ahead of those who seek to claim it.",
				2: "An adventurer got me out of prison, and I'll never go back.",
				3: "My allies turned on me and left me for dead. Adventuring will help me build a new life.",
				4: "Adventuring is easy coin, nothing more.",
				5: "Adventuring tests the limits of my skills, something crime hasn't done in years.",
				6: "I've done too many dark things to ever sleep well, but adventuring might help me make amends.",
				7: "It's time to dismantle the criminal guild I helped create, and adventuring will give me the power to do so.",
				8: "A dogged constable is after me, but even they won't go where adventuring will take me.",
			},
		},
	},
	"homesteader": {
		Name:             "Homesteader",
		BackgroundSource: "Players Guide, pg 121",
		Description: "You forged a livelihood in the places between civilization and the unknown hinterlands. " +
			"The demands of frontier life calloused you, but you understand the wilderness and your place " +
			"in it. Perhaps you were a weather-beaten frontiersman done with the lumber camps behind, a " +
			"hermit who wished to quit your seclusion, or a young hunter ready to test your mettle on " +
			"more dangerous prey. No matter, you forged your own path in a harsh wilderness, and those " +
			"skills will only help you forge ahead.",
		SkillProficiencies: []string{"Survival"},
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 1,
				Options:        []string{"Athletics", "Animal Handling", "Intimidation"},
			},
		},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"tools": {
				NumberToSelect: 1,
				Options:        []string{"herbalism tools", "navigator tools"},
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "hunting trap", Quantity: 1},
			{Name: "fishing tackle", Quantity: 1},
			{Name: "skinning knife", Quantity: 1},
			{Name: "hammock", Quantity: 1},
			{Name: "clothes traveler heavy", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 8},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"aware", "dungeoneer", "far traveler"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "Adventuring will give me new challenges to overcome without the expectation of settling down.",
				2: "I've been alone for too long. Adventuring will allow me to find companionship.",
				3: "My name is all I have, and adventuring will help it grow into something to be proud of.",
				4: "Adventuring will take me to exotic places where I may start my next, or last, expedition.",
				5: "Hopefully adventuring will give me enough coin to buy back the camp that was stolen from me.",
				6: "Settling down didn't work for me, so adventuring is how I find thrills.",
				7: "I accompany travelers on their adventure for pay—promises, gold, or favors.",
				8: "I've yet to find anyone as skilled and reliable in the wilds as I am, but adventuring may change that.",
			},
		},
	},
	"maker": {
		Name:             "Maker",
		BackgroundSource: "Players Guide, pg 122",
		Description: "You pursued a unique, often profitable craft and became an expert. Those with an eye " +
			"for quality might seek your work out among hundreds of other crafters. Perhaps you were the " +
			"disgraced scion of an illustrious family of jewelers, a famous swordsmith ready to test your " +
			"finest work, or a toy maker who manufactured clockwork monstrosities. No matter what your " +
			"artform, you worked so fervently that it stays a part of you wherever you go.",
		SkillProficiencies: []string{"Investigation"},
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 1,
				Options:        []string{"History", "Performance", "Sleight of Hand"},
			},
		},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"tools": {
				NumberToSelect: 1,
				Options:        static_data.GetToolsetNames(),
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "wax seal", Quantity: 1},
			{Name: "clothes traveler", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 10},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"artillerist", "school specialization", "trade skills"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "I seek inspiration so divine or perilous that only adventuring may provide it.",
				2: "Adventuring allows me to test my creations to the fullest.",
				3: "Adventuring aids me in discovering rare and otherwise unknown ingredients.",
				4: "Jealousy drove my peers to chase me from my workshop, but adventuring may secure my fortune once again.",
				5: "Adventuring is how I make the coin required to fund my artifice to its fullest.",
				6: "I have yet to find an equal in my craft, and I hope that in adventuring one will cross my path.",
				7: "Adventuring is the only way to grow my fame, as my craft is too unique or obscure for common minds.",
				8: "Adventuring is the only way I might find someone worthy of possessing my greatest work.",
			},
		},
	},
	"outcast": {
		Name:             "Outcast",
		BackgroundSource: "Players Guide, pg 122",
		Description: "You spent your life surviving on scraps and taking what you could. Living on the " +
			"streets sometimes left you on the wrong side of the law, but you were instilled with skills " +
			"to survive, overcome, and prosper. Perhaps you were an urchin chased from your stomping " +
			"grounds, a pickpocket who tried to make ends meet, or a bandit who left the life, wanting " +
			"to make amends. Whatever your circumstances, the thrills and misfortunes of life outside " +
			"polite society will never leave you.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"Deception", "Insight", "Sleight of Hand", "Stealth"},
			},
		},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"games": {
				NumberToSelect: 1,
				Options:        []string{"game set"},
			},
			"tools": {
				NumberToSelect: 1,
				Options:        []string{"charlatan tools", "herbalism tools", "thieves tools"},
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "cloak dark", Quantity: 1},
			{Name: "clothes common dark", Quantity: 1},
			{Name: "silver coin", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 10},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"aware", "opportunist", "quick"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "Adventuring is a way to stay ahead of the law, I hope.",
				2: "Adventuring is how I'll finally earn (or seize) my fortune.",
				3: "I will amass power and influence by adventuring before I return home.",
				4: "Adventuring is how I'll make amends for a life of wrongdoing.",
				5: "I can master my skills through adventuring without fear of ending up in a cell.",
				6: "Adventuring will give me the clout to make a name for myself that will be feared and respected.",
				7: "I'll track down the person who ruined my life while I'm adventuring.",
				8: "Adventuring is a way to find a crew I can trust.",
			},
		},
	},
	"rustic": {
		Name:             "Rustic",
		BackgroundSource: "Players Guide, pg 123",
		Description: "You spent most of your life as no one of consequence. Years of hard work gave you " +
			"an unshakeable resolve, but your past is no mystery and affords you no grand understanding " +
			"of the world. Perhaps you were the blacksmith's child who preferred to wear the armor, a " +
			"shepherd who watched her flock devoured by ogres, or an elderly dwarf miner who wanted to " +
			"see the world before the end. Wherever you come from, whoever you were, even a perilous " +
			"future seems better than the doldrums of your past.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"Athletics", "Acrobatics", "Investigation", "Medicine"},
			},
		},
		AdditionalProficiencies: []string{"land vehicles"},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"equipment": {
				NumberToSelect: 1,
				Options:        []string{"martial weapon", "musical instrument", "tool", "armor"},
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "backpack", Quantity: 1},
			{Name: "bedroll", Quantity: 1},
			{Name: "blanket", Quantity: 1},
			{Name: "candle", Quantity: 3},
			{Name: "clothes traveler", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{SilverPieces: 20},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"comrade", "hand to hand", "physical fortitude"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "Adventuring gives me thrills I never experienced back home.",
				2: "Adventuring supplies coin that will secure a better future—if not for me, for my family.",
				3: "I'd rather risk my life adventuring than waste it in obscurity.",
				4: "Adventuring will give me such stories to tell around the fire when it's time to settle down again.",
				5: "I can't face my friends after what I did—not until I make a name for myself through adventuring.",
				6: "Maybe adventuring can teach me the skills I need to become a noble.",
				7: "I was blamed for the ill fate that befell my home. Maybe by adventuring, I can make things right.",
				8: "An adventurer saved my life, and I won't rest until I do the same for others.",
			},
		},
	},
	"scholar": {
		Name:             "Scholar",
		BackgroundSource: "Players Guide, pg 123",
		Description: "You spent years researching a branch of study. Time spent in academic pursuits honed " +
			"your mind, allowing you to view the world through an intellectual lens afforded to few. " +
			"Perhaps you were only recently a student eager to learn outside the classroom, a teacher " +
			"who retired but wasn't ready to stop hands-on learning, or a discredited researcher expelled " +
			"but driven to prove your theories. Regardless, your way has always been lit by your keen " +
			"mind, and you retain a desire to know more.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"Arcana", "History", "Nature", "Religion"},
			},
		},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"languages": {
				NumberToSelect: 2,
				Options:        LanguageNames(),
			},
			"tools": {
				NumberToSelect: 1,
				Options:        static_data.GetToolsetNames(),
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "ink", Quantity: 1},
			{Name: "quill", Quantity: 1},
			{Name: "knife small", Quantity: 1},
			{Name: "reference book", Quantity: 1},
			{Name: "clothes common", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 10},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"polyglot", "ritualist", "school specialization"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "The coin I need for my research comes from adventuring.",
				2: "Adventuring provides valuable field experience relevant to my study.",
				3: "I will prove those fools wrong with my discoveries made through adventuring.",
				4: "Adventuring pays the bills until I can prove the validity of my theories.",
				5: "I can uncover lost or forbidden knowledge by adventuring for it. No institution can provide that!",
				6: "Adventuring is the best way to collect data for my patron or employer.",
				7: "Adventuring will lead me to the answers I desperately seek.",
				8: "Adventuring is a way to escape a life of academia I never wanted.",
			},
		},
	},
	"soldier": {
		Name:             "Soldier",
		BackgroundSource: "Players Guide, pg 124",
		Description: "You spent a significant amount of time risking your life to defend others. You survived " +
			"through rigorous training, discipline, and sacrificing comforts that most people take for " +
			"granted. Perhaps you were a veteran who washed out, a deserter who ran from the atrocities " +
			"of war, or a fresh-faced patriot who went looking for new ways to fight for your cause. " +
			"Whatever course you took, you remain forever changed having borne the weight of duty.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"Animal Handling", "Athletics", "Medicine", "Survival"},
			},
		},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"tools": {
				NumberToSelect: 1,
				Options:        static_data.GetToolsetNames(),
			},
			"vehicles": {
				NumberToSelect: 1,
				Options:        []string{"land vehicles"},
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "symbol of rank", Quantity: 1},
			{Name: "mess kit", Quantity: 1},
			{Name: "playing cards", Quantity: 1},
			{Name: "clothes common", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 10},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"combat casting", "combat conditioning", "field medic"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "After a dishonorable discharge, adventuring is the way I make a living.",
				2: "Adventuring is a way to continue fighting, even though the war is over.",
				3: "Adventuring is a way to keep protecting others, since those I used to protect are gone.",
				4: "Adventuring lets me use the skills I learned without having to give my life to the military.",
				5: "I perform special missions for those I serve when I go adventuring.",
				6: "When I go adventuring, I take justice into my own hands without concern for policy or politics.",
				7: "Adventuring is the return to action I've craved since my retirement.",
				8: "Adventuring is a way to keep my skills sharp before I can return to duty.",
			},
		},
	},
}
