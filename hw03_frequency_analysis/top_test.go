package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var quote = `“Our opportunities to give of ourselves are indeed limitless, but they are also perishable.
	There are hearts to gladden. There are kind words to say. There are gifts to be given. There are deeds to be done.
	There are souls to be saved. As we remember that “when ye are in the service of your fellow beings ye are only
	in the service of your God,” (Mosiah 2:17) we will not find ourselves in the unenviable position of Jacob Marley’s
	ghost, who spoke to Ebenezer Scrooge in Charles Dickens’s immortal "Christmas Carol." Marley spoke sadly
	of opportunities lost. Said he: 'Not to know that any Christian spirit working kindly in its little sphere,
	whatever it may be, will find its mortal life too short for its vast means of usefulness. Not to know that no
	space of regret can make amends for one life’s opportunity misused! Yet such was I! Oh! such was I!
	'Marley added: 'Why did I walk through crowds of fellow-beings with my eyes turned down, and never raise them
	to that blessed Star which led the Wise Men to a poor abode? Were there no poor homes to which its light would
	have conducted me!'Fortunately, as we know, Ebenezer Scrooge changed his life for the better.
	I love his line, 'I am not the man I was.'Why is Dickens’ "Christmas Carol" so popular? Why is it ever new?
	I personally feel it is inspired of God. It brings out the best within human nature. It gives hope.
	It motivates change. We can turn from the paths which would lead us down and, with a song in our hearts,
	follow a star and walk toward the light. We can quicken our step, bolster our courage, and bask in the
	sunlight of truth. We can hear more clearly the laughter of little children. We can dry the tear of the weeping.
	We can comfort the dying by sharing the promise of eternal life. If we lift one weary hand which hangs down,
	if we bring peace to one struggling soul, if we give as did the Master, we can —by showing the way— become
	a guiding star for some lost mariner.” - Thomas S. Monson`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		assert.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			assert.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			assert.ElementsMatch(t, expected, Top10(text))
		}
	})

	t.Run("latin positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			assert.Contains(t, Top10(quote), "i")
		} else {
			assert.NotContains(t, Top10(quote), "i")
		}
	})
}
