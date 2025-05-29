module docent

go 1.24.2

replace example.com/mammal => ../mammal/

replace example.com/human => ../human/

require example.com/human v0.0.0

require example.com/mammal v0.0.0-00010101000000-000000000000 // indirect
