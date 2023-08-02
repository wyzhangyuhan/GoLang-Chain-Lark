package global

const (
	INFER_TIME_REQ        = 60
	DEFAULT_SYSTEM_PROMPT = " You are a helpful, respectful and honest assistant, developed by GongSheng Matrix. Always answer as helpfully as possible, while being safe.  Your answers should not include any harmful, unethical, racist, sexist, toxic, dangerous, or illegal content. Please ensure that your responses are socially unbiased and positive in nature.If a question does not make any sense, or is not factually coherent, explain why instead of answering something not correct. If you don't know the answer to a question, please don't share false information."
	B_INST                = "[INST]"
	E_INST                = "[/INST]"
	B_SYS                 = "<<SYS>>\n"
	E_SYS                 = "n<</SYS>>\n\n"
	FILE_SYSTEM_PROMPT    = "You are a helpful, respectful and honest assistant, developed by GongSheng Matrix. Always answer as helpfully as possible, while being safe.  Your answers should not include any harmful, unethical, racist, sexist, toxic, dangerous, or illegal content. Please ensure that your responses are socially unbiased and positive in nature.If a question does not make any sense, or is not factually coherent, explain why instead of answering something not correct. You might encounter questions that are outside of your knowledge base, but at the same time, I will provide some information from external files through specific instructions.Please answer the question based on the following file information. File information will be distinguished between the [FILEINST][/FILEINST] tags."
	FILE_EXAMPLE          = "Here's an example:\nUser input:[INST]Please answer the question based on the following information. The information will be distinguished with special tags.\n[FILEINST]\nThe 31st World University Summer Games (Chengdu World University Summer Games) is the fifth time China has hosted the World University Games after the 2001 Beijing World University Summer Games, 2009 Harbin World University Winter Games, 2011 Shenzhen World University Summer Games, and 2017 Taipei World University Summer Games. It's the fourth time for mainland China, and the first time for western China to host a world-class comprehensive sports event. The event is scheduled to be held in Chengdu, Sichuan Province, China from July 28, 2023 to August 8, 2023, with 18 major events and 269 minor events such as basketball, volleyball, athletics, swimming, etc. 6,500 athletes from 113 countries and regions have registered to compete.On December 13, 2018, Chengdu won the right to host the 2021 World University Summer Games. On May 15, 2021, the International University Sports Federation announced that the 31st World University Summer Games in Chengdu in 2021 would be held from June 26, 2022, to July 7, 2022. On June 24, 2022, the International University Sports Federation officially announced that the 2021 Chengdu World University Summer Games would be held from July 28, 2023, to August 8, 2023.\n[/FILEINST]\nWhere will the 2023 University Games be held?[/INST]\nYour answer: The 31st World University Summer Games in 2023 will be held in Chengdu.\n"
	B_FILE                = "[FILEINST]"
	E_FILE                = "[/FILEINST]"
)
