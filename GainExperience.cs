/*
Gain Experience Function

Created by: JarekQ Aloisio
Created on: 08/08/2019
Last Modified: 08/19/2019
Language: C-Sharp (C#)
*/
// ---------------
/*<summary> 
This function returns an integer amount of the total experience gained for every challenge faced.
Success is defined by learning a new thing while you do YourBest() function.
Failure is NOT BAD, it adds to the learning heuristic of the YourBest() function.
Each time you Do Your Best, you are constantly learning and achieving.
Your experience is the summation of all challenges faced, every success gained, and every failure experienced.
Experience is the sum all parts of doing your best, that ultimately defines who you are!
</summary>*/
// ---------------
// Pre-Condition: This function takes in an Integer value for total amount of Challenges.
// ---------------
// Post-Condition: Returns an Integer value for the total amount of Experience gained.
// ---------------
public int GainExperience(int challenges_amount)
{
	int failure = 0;
	int success = 0;
	bool learnedNewThing = false;
	int experience = new int();

	for(experience = 0; experience < challenges_amount; experience++)
	{
		do{
			learnedNewThing = YourBest(success,failure);
			
			if(learnedNewThing == true)
			{
				success++;
			}
			else
			{
				failure++;
			}
		}while(!learnedNewThing);
		
		learnedNewThing=false;
	}

	experience += (success + failure);

	return experience;
}