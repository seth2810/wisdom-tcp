package quotes

import (
	"math/rand/v2"
)

var quotes = []string{
	`"Walking on water and developing software from a specification are easy if both are frozen." - Edward V Berard`,
	`"Computer science education cannot make anybody an expert programmer any more than studying brushes and pigment can make somebody an expert painter." - Eric S. Raymond`,
	`"Talk is cheap. Show me the code." - Linus Torvalds`,
	`"In God we trust. All others must bring data." - W. Edwards Demming`,
	`"In theory, theory and practice are the same. In practice, they're not." - Yoggi Berra`,
	`"You can't have great software without a great team, and most software teams behave like dysfunctional families." - Jim McCarthy`,
	`"The best programmers are not marginally better than merely good ones. They are an order-of-magnitude better, measured by whatever standard: conceptual creativity, speed, ingenuity of design, or problem-solving ability." - Randall E. Stross`,
	`"People think that computer science is the art of geniuses but the actual reality is the opposite, just many people doing things that build on each other, like a wall of mini stones." - Donald Knuth`,
	`"Most of you are familiar with the virtues of a programmer. There are three, of course: laziness, impatience, and hubris." - Larry Wall`,
	`"Most good programmers do programming not because they expect to get paid or get adulation by the public, but because it is fun to program." - Linus Torvalds`,
	`"Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live." - Martin Golding`,
	`"Nine people can't make a baby in a month." – Fred Brooks`,
	`"Any fool can write code that a computer can understand. Good programmers write code that humans can understand." – Martin Fowler`,
	`"Programming is like sex. One mistake and you have to support it for the rest of your life." – Michael Sinz`,
	`"If builders built buildings the way programmers wrote programs, then the first woodpecker that came along would destroy civilization." – Gerald Weinberg`,
	`"Debugging is twice as hard as writing the code in the first place. Therefore, if you write the code as cleverly as possible, you are, by definition, not smart enough to debug it." — Brian W. Kernighan`,
	`"Some people, when confronted with a problem, think "I know, I'll use regular expressions." Now they have two problems." — Jamie Zawinski`,
	`"You can use an eraser on the drafting table or a sledgehammer on the construction site" - Frank Lloyd Wright`,
	`"The first 90% of the code accounts for the first 90% of the development time. The remaining 10% of the code accounts for the other 90% of the development time." — Tom Cargill`,
	`"There are only two kinds of languages: the ones people complain about and the ones nobody uses." — Bjarne Stroustrup`,
	`"It's all talk until the code runs." — Ward Cunningham`,
	`"Given enough eyeballs, all bugs are shallow." - Linus Torvalds`,
	`"A clever person solves a problem. A wise person avoids it." — Albert Einstein`,
	`"Being a good software engineer is 3% talent, 97% not being distracted by the internet." — Unknown`,
	`"Any sufficiently advanced technology is indistinguishable from magic." - Arthur C. Clarke`,
	`"A ship in port is safe, but that's not what ships are built for." - Grace Hopper`,
	`"The imposter syndrome is real. Luckily, it goes away." - Kimber Lockhart`,
	`"Know how to learn. Then, want to learn." - Katherine Johnson`,
	`"Never trust a computer you can't throw out a window." - Steve Wozniak`,
	`"When in doubt, use brute force." - Ken Thompson`,
	`"Once a new technology starts rolling, if you're not part of the steamroller, you're part of the road." - Stewart Brand`,
	`"The most disastrous thing that you can ever learn is your first programming language." - Alan Kay`,
	`"Optimism is an occupational hazard of programming. Feedback is the treatment." - Kent Beck`,
	`"The function of good software is to make the complex appear to be simple." - Grady Booch`,
	`"I'm not a great programmer. I'm just a good programmer with great habits." ― Kent Beck`,
}

func GetRandomQuote() []byte {
	idx := rand.IntN(len(quotes))

	return []byte(quotes[idx])
}
