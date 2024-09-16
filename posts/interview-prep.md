# Interview prep tips

I have done 300+ interviews at this point and helped 15+ people prepare for those. This page sums up the common topics that I regularly see in interviews and interview prep.

## Strategy

### Strategy during your preparation

&gt; "Consistency beats intensity"

It's better to aim for 1-2 exercise per day - at minima 2-3 times per week - rather than do 10 exercises on Saturday/Sunday. The goal is to do a bit each day. This approach helps to cement your understanding.

&gt; Your goal is to **drill in some good habits and** solve the problem

You're read correctly, it's both!

Solving a problem is only half of the battle in an interview. In fact, we expect you to be able to understand the problem, communicate your algorithm, implement it, find your own bugs, talk about its big-O notations, optimize, ... This is why the methodologies below should be central to your practice. They hit a lot of skills that interviewers look for in an interview.

&gt; Rock solid Computer Science foundation

If you don't know all your data-structures and algorithms (DSA), you **WILL FAIL IN YOUR INTERVIEWS**. There is no way around it, you must know all those pesky DSA that you studied in your Computer Science class(es). Bonus points if you are able to implement data-structures from scratch.

There is no way around it: you must learn by heart the big-O notations **for all DS and for all operations** (insert, replace, delete, lookup). The most efficient way to do this is [spaced repetition](https://en.wikipedia.org/wiki/Spaced_repetition), either using physical flash cards or an app, like [Anky](https://apps.ankiweb.net/).

&gt; Implement, optimize, make pretty (in this order)

Your goal for each exercise should be to __implement a solution__. It doesn't matter if the solution is O(N^2) or the code is ugly. We hire you to solve problems by writing working code!

Once you can implement an exercise with a brute-force solution, you can look at implementing optimized solutions. If you can't implement a solution, then it's not really useful in an interview.

Once you can solve an exercise with an optimized solution, you can look at making it pretty by e.g.:

* adding some helper functions ([check this StackOverflow thread if you don't know what this is](https://stackoverflow.com/questions/42837402/role-of-helper-functions))
* removing [magic constants](https://en.wikipedia.org/wiki/Magic_number_(programming)#Unnamed_numerical_constants)
* adding classes/objects
* adding documentation (e.g. JavaDoc, PyDoc or equivalent)

Note: This list is non-exhaustive.

### Strategy during the interview

&gt; Your goal in the interview should be to code something

Ideally this should be some optimized algorithm... but if you're struggling to find such an optimized/elegant solution, offer to implement the brute force solution and revisit the algorithm later. Listen to any guidance from your interviewer.

&gt; Listen to your interviewer

Each interviewer is different so the best tip is to listen to what they tell you. Make sure you **understand what they meant**, which may be different from what they said unfortunately.

An interviewer may interrupt you to ask side questions or even move to a different question. This is normal and expected. Their questions/directives reflect what skill(s) they want to evaluate.

If an interviewer gives you an instruction, make sure to acknowledge, clarify and ultimately follow it... or be prepared to explain why you didn't (it's possible the interviewer made a mistake).

Also a good interviewer will step in with some hint if you're struggling - usually after 1-2 minutes. Make sure to listen to what they say as they are trying to help you!

&gt; If you are struggling, you're not necessarily doing bad

The goal of an interview is to find your level in some key software engineering skills. This means that a good interviewer will constantly be pushing you. The corollary is that there is some expected amount of struggle that should happen.

Also if an interviewer changes the problem because you're struggling too much, it's **not the end of the interview**. You're given a chance to recover and you should take it. I have seen some candidate struggle with a specific domain (e.g. strings) but do really well in another (e.g. hash tables). This is just an indication of an area for improvement, something normal for every person, that may not be a blocker for hiring.

## Methodology for solving algorithm questions

This is my #1 feedback in mock interviews as it's something that is not taught to new engineers. It have given it to **almost every single mock interview** I have conducted.

You **NEED a process for approaching coding/algorithm questions**. If you think you don't, think again: I have given this advice many times but forgot to follow it, which led to a spectacular crash-and-burn in one specific interview (not my proudest moment!).

There are 2 main methodology that I know of:

* Cracking the code interview
* UMPIRE

Both are fairly equivalent so it doesn't matter which one you choose. What matters is that you pick up one and stick to it during your interview prep.

Drilling the chosen methodology until it becomes a habit is mandatory. Due to the pressure in an interview, you will fall back to your level of preparation. This means that moving forward **EVERY LeetCode problem should follow the chosen methodology**.

Let's dive into each one.

### Cracking the code interview

If you have [the book](https://www.crackingthecodinginterview.com/), you can find the methodology on _page 62_. Page 60 through 66 explain how to apply it so those are good pages to read first. I will just summarize the 7 steps below, it's **strongly recommended that you get access to the book**.

Note: While the book is close to 10 years old at this point, most of its advices are universal so it is a good investment in your career. Also your local libraries may have it.

This is a 7 steps process:

1. Listen (understand the problem, ask some clarifying questions)
2. Example (run an example with 3-5 "elements" (nodes in a tree/graph, items in an array, ...) to find out an algorith to solve it)
3. Brute force (quickly outline the brute force algorithm)
4. Optimize (try to find a better algorithm. If you can't find any, go to step 5.)
5. Walk through your algorithm (test your algorithm against the example from step 2.)
6. Code
7. Test (test your code against the example from step 2.)

Here are some resources on it:
* [YT series showing how to solve problems using the methodology](https://www.youtube.com/playlist?list=PLI1t_8YX-ApvFsH-DaFmAmdJboAnbg08P)

### UMPIRE

This method was developed by [CodePath.org](https://www.codepath.org/).

UMPIRE stands for:

* Understand (make sure you understand the problem, ask some clarifying questions)
* Match (does the problem match any known algorithm(s)? can you apply them here too?)
* Plan (describe your algorithm, write some pseudocode)
* Implement (code!)
* Review (check your code against an example with 3-5 "elements" (nodes in tree/graph, items in array, ...))
* Evaluate (big-O runtime/space, outline any potential optimization or improvements to the code)

Here are some resources on UMPIRE:

* [CodePath seminal guide with an example)](https://guides.codepath.com/compsci/UMPIRE-Interview-Strategy)
* [Another good guide](https://dahadaller.github.io/umpire/)
* [Video about solving a linked list problem with UMPIRE](https://www.youtube.com/watch?v=W6V7MLE_5X4)
* [Intro video on solving a simple problem using UMPIRE](https://www.youtube.com/watch?v=wBRYtW-TJq8) (this is the first part of series on using UMPIRE on harder and harder problems that is worth checking out)
* [UMPIRE in interviews/on-the-job](https://medium.com/@rmorenocesar/codepaths-umpire-method-in-the-wild-a884c05b96fc) ([ADVANCED] this article outlines the steps in UMPIRE and link those to what to do in an interview or once you've landed a job)

## Resources

### Study plans

* [Coding Interview study plan](https://www.techinterviewhandbook.org/coding-interview-study-plan/)
* [List of LeetCode problems based on the available time](https://www.techinterviewhandbook.org/grind75) [good if you don't know where to start]
* [LeetCode discussion about a study plan](https://leetcode.com/discuss/general-discussion/460599/blind-75-leetcode-questions)
* [Curated list of LeetCode problems based on time before interviewing](https://jeremyaguilon.me/blog/ranking_interview_questions_by_cram_score)
