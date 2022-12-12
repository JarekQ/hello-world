/*
Great Things Function

Created by: JarekQ Aloisio
Created on: 01/17/2019
Last Modified: 10/01/2019
Language: C-Sharp (C#)
*/
//------------------------
/*<summary>
This simple program defines two main things:
1. The state of the user at work.
2. The function the user performs when present at work.
The program ends when the user leaves work after performing the GreatThings() function.
</summary>*/

using System;
using Q.Code;

class Main
{
	static void Main(string[] args) 
	{
		bool yourPresence = user.GetState();
		
		do{
			GreatThings();
		}while(yourPresence == true);
		
		return 0;
	}
}