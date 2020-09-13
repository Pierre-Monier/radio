module radio

go 1.13

replace radio/rutils => ../rutils

replace radio/rweb => ../rweb

require (
	radio/rutils v0.0.0-00010101000000-000000000000
	radio/rweb v0.0.0-00010101000000-000000000000
)
