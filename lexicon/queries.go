package lexicon

const GetQuery string = `
select 
	l.id,
	u.lexema,
	u.canonica,
	coalesce(l.categoria, ''),
	coalesce(l.forma, ''),
	coalesce(l.tempo, ''),
	coalesce(l.tipo, ''),
	coalesce(l.valor, ''),
	coalesce(tl.traco, '')
from unidade_lexica u
inner join lexico l on (l.unidade_lexica = u.id)
left join traco_lexico tl on (tl.lexico = l.id)
where u.lexema = '%s'`

const PostQuery string = `
with lexema as (
	select id 
	from palavra
	where palavra.id = '%s'
),
canonica as (
	select id 
	from palavra
	where palavra.id = '%s'
),
ul as (
	SELECT id
	FROM unidade_lexica
	where 
		unidade_lexica.lexema = (select id from lexema)
		and unidade_lexica.canonica = '%s'
)
INSERT INTO lexico
(unidade_lexica, categoria, tipo, forma, tempo, valor)
VALUES((select id from ul), '%s', %s, %s, %s, %s) --could be NULL for tipo, forma...
RETURNING id`

const FlexaoQuery string = `
INSERT INTO traco_lexico
(lexico, traco)
VALUES(%s, '%s');`

const DelLexiconQuery string = `
DELETE FROM public.lexico
WHERE id=%s;`

//Quebra a ConCom em duas palavras
const ConComQuery string = `
select 
	distinct on (up.lexema)
	--lp.categoria as "categoria",
	up.lexema as "palavra"
	--coalesce(lp.tipo, '') --to avoid nil exception
from unidade_lexica u
inner join lexico l on (l.unidade_lexica = u.id)
left join contracao c on (c.id = l.id)
left join lexico lp on (lp.id = c.prefixo or lp.id = c.sufixo)
left join unidade_lexica up on (up.id = lp.unidade_lexica)
where u.lexema = '%s'
order by 
	up.lexema, c.id asc`
