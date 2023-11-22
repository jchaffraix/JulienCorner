# Productivity

## Productivity vs efficiency

<!--TODO: Insert a vector diagram of velocity vs speed.-->

Productivity is a raw measure of output. In a factory, productivity would be the absolute number of widgets that is made in a day. In productivity systems, that would be the number of tasks tackled.

Efficiency (or effectiveness) is the measure of productivity towards a goal. The distinction is important as anything done that is in service of a goal is effectively wasted.

Someone can be very productive (gets a lot done) but very ineffective (does the wrong thing), conversely someone can be very effective but very unproductive (the low output is 100% placed towards a goal).

In real projects, the distinction between productivity and efficiency is harder to draw. This is because there are inherent tensions between different goals. For example, learning a new way of solving a task makes you slower at doing tasks (lower productivity). However this tends to increase the productivity in the long run (see the [leverage section](#leverage)). Another example is adding some protection to a backend (increasing reliability) will increase its latency (worsening the performance).

There is however one category of tasks that are clearly unneeded, usually called "waste".

### Waste

Waste is central to the [lean thinking](https://en.wikipedia.org/wiki/Lean_thinking) framework. There is even have a word for it in the Toyota Production System (the origin of the lean movement): [muda](https://en.wikipedia.org/wiki/Muda_(Japanese_term)).

Lean thinking focuses on the value stream: providing value to the customer in an efficient way. As such, lean thinking spends a large amount of time looking for waste.

Lean thinking has 2 main incarnations in the software world:
1. Lean in software development (focusing on efficient software producing)
2. Lean in product development (focusing on the efficient determination of the next product/feature to build).

Both define different kind of waste but we can group them in a few categories:

**Unneeded tasks**

Anything that is done but doesn't benefit the goal has an efficiency of 0. The collary is that you should spend as little as needed to get the tasks done.

Something that is an art in the software work is deciding the level of quality for a feature. A start-up before product/market fit with few users benefits from prototypes (low quality features), prefering a high iteration rate. This is because they are exploring the problem space and thus want to limit the waste while doing so. Conversely, an established company would prefer higher quality features to ensure better reliability for their existing customers. Work to increase the quality above what is expected is unneeded.

A side note on quality, there is wiggling room in the level of quality required. I am not suggesting to ship half-working software as that would be counter-productive. However there is a strong value in shipping software and getting users' feedback. I associate with the [software crafsmanship movement](https://manifesto.softwarecraftsmanship.org/#/en/reading) in that we should ship high quality software, but recognize that there are times when shipping lower quality to limit the waste is a worthy endeavor.

There are other tasks besides quality where there are some variance on the output that are worth talking about.

The first one is scoping. Scoping is the art of determining how long a project will take. Most engineers use some heuristics for doing that (estimate the lengh then double-triple your original estimate). However something to recognize is that there are times where doing rough scoping (this project will take 2 weeks -/+ 2 weeks) is a good enough estimation. Doing a fine grained one (the project will take 2.75 weeks) means spending more time doing it, which is wasted if we decide not to proceed with the project.

Another such task is coming up with proposals. As you get more experience, it is expected that you would come up with new projects. However before there is agreement on whether it is worth doing a specific project, you need to spend just enough time to convince yourself and your team that it is worth doing. Spending more time than this is effectively wasted.

**Toil**

The [Google SRE handbook](https://sre.google/sre-book/eliminating-toil/) defines toil as:

> Toil is the kind of work tied to running a production service that tends to be manual, repetitive, automatable, tactical, devoid of enduring value, and that scales linearly as a service grows.

Toil is a special kind of unneeded task. It is a task that can be easily done by a computer.

<a id="leverage"></a>
## Leverage

A good way to increase your personal productivity is to use leverage.

Effort applied to a task is not the same as impact towards solving it.

Leverage is the idea that some tasks multiply your productivity or effectiveness, akin to how a lever multiplies a force.

If you've ever heard the the motto "work smarter, not harder", this is leverage &#x1F600;

One of the good resources on leverage in software is the [Effective Engineer book](https://www.effectiveengineer.com/). The book defines 5 areas of
high leverage:

1. Optimise for learning
2. Invest in iteration speed
3. Validate your ideas aggressively and iteratively
4. Minimise operational burden
5. Building a great engineering culture

Those areas translate into several activities you can do. For time's sake, I will touch on the 2 I find most important, see the video at the end for more.

The first one is tooling. Tooling is a form of automation that increases your iteration speed and free time for other tasks. Investing in tooling early and often will strongly increase your speed.

The second one is hiring. Hiring someone will boost your productivity as you can now delegate the task to the person. It is not a 2x but it can get pretty close.

## Focus

Focus is central to a efficiency. The reason is that you can't make progress if you have to do 20 tasks. This is why lean software development limits the number of work-in-progress to focus on a few tasks.

Further more, the cost of switching between tasks is completely waste. For creative tasks, it can take 30-60m to get into a task that can be wasted due to an interruption.

There are a lot of literature on the topics that I have listed at the end of the article but to summarize on those.

To help focus, only work on a single task at a time, until you complete it. Once it is completed, switch to the next tasks.

For maximum efficiency, sort your task by efficiency and start with your most important tasks.

To foster focus, ask a clarification like "What's the ONE Thing I can do such that by doing it everything else will be easier or unnecessary?" (taken from the [One Thing](https://en.wikipedia.org/wiki/The_One_Thing_(book))).

## Further resources

2 articles by Mark Manson about personal productivity for creative people. A strongly recommended read:
1. [https://markmanson.net/how-to-be-more-productive](https://markmanson.net/how-to-be-more-productive)
2. [https://markmanson.net/principles-of-productivity](https://markmanson.net/principles-of-productivity)


Leverage:
* [The talk@Google Youtube video about the Effective Engineer](https://www.youtube.com/watch?v=BnIz7H5ruy0).


[My book recommendations on productivity](/pages/books.html#productivity)
