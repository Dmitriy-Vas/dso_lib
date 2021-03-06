package plugins

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/dso_lib/packets/incoming"
	"github.com/Dmitriy-Vas/dso_lib/packets/outgoing"
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave_buffer"
	"log"
	"sync"
)

// AllAchievements claims all steam achievements after receiving ClientRevisionPacket
func AllAchievements(proxy *wave.Proxy) {
	var once = sync.Once{}
	var buf buffer.PacketBuffer = nil
	proxy.HookPacket(int64(dso_lib.OutClientRevision), true, func(conn *wave.Conn, packet wave.Packet) {
		if p := packet.(*outgoing.ClientRevisionPacket); !p.IsSteam {
			log.Printf("Please run game from the Steam to claim achievements")
			return
		}

		once.Do(func() {
			buf = conn.Buffer().Clone()
		})

		done := make(chan error)
		defer close(done)
		go func(done chan error, conn *wave.Conn) {
			for {
				err, ok := <-done
				if ok {
					if err != nil {
						conn.Close(err)
						break
					}
				} else {
					log.Printf("All achievements claimed")
					break
				}
			}
		}(done, conn)

		p := &incoming.SteamAchievementPacket{
			ID:              int64(dso_lib.IncSteamAchievement),
			Send:            true,
			Achievement:     dso_lib.ACHIEVEMENT_AW_QUESTCOMPLETE,
			AchievementStat: dso_lib.ACHIEVEMENT_STAT_QUESTCOUNT,
		}
		for i := 0; i <= 100; i++ {
			done <- conn.SendPacket(buf, p, false)
		}

		p.Achievement = dso_lib.ACHIEVEMENT_AW_GUMBYKILL100K
		p.AchievementStat = dso_lib.ACHIEVEMENT_STAT_GUMBYCOUNT
		for i := 0; i <= 100000; i++ {
			done <- conn.SendPacket(buf, p, false)
		}

		p.Achievement = dso_lib.ACHIEVEMENT_AW_MONSTERKILL
		p.AchievementStat = dso_lib.ACHIEVEMENT_STAT_KILLCOUNT
		for i := 0; i <= 100000; i++ {
			done <- conn.SendPacket(buf, p, false)
		}

		p.Achievement = dso_lib.ACHIEVEMENT_AW_PLAYERKILLER
		p.AchievementStat = ""
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_LICHCAPE
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_HALLOWEENTOWER
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_DEATH
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_KILL_LUCAS
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_ANNIVERSARY
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_CREATECHAR
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_BUYPREMIUMITEM
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_ENHANCEITEM
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_FIRSTMOUNT
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_REACHESIA
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_REACHGOLDUM
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_REACHTERION
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_REACHROKKUNTAVERN
		done <- conn.SendPacket(buf, p, false)

		p.Achievement = dso_lib.ACHIEVEMENT_AW_REACHCEMETERY
		done <- conn.SendPacket(buf, p, false)
	})
}
