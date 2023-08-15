import { Component, OnInit } from '@angular/core';
import { DateTime } from 'luxon';
import { Match } from '../../model/match';
import { EloService } from '../../services/elo.service';
import { MessageBoxService } from '../../services/message-box.service';

@Component({
  selector: 'app-matches',
  templateUrl: './matches.component.html',
  styleUrls: ['./matches.component.sass']
})
export class MatchesComponent implements OnInit {
  public displayedColumns = ["dateTime", "player1", "score", "player2", "delete"];
  public dataSource: Match[] = [];
  public isLoaded = false;
  public newDate?: Date;
  public newTime?: string;
  public newPlayer1 = new PlayerScoreViewModel();
  public newPlayer2 = new PlayerScoreViewModel();

  constructor(private eloService: EloService,
    private messageBoxService: MessageBoxService) {

  }

  public async ngOnInit(): Promise<void> {
    await this.refresh();
  }

  private async refresh(): Promise<void> {
    this.isLoaded = false;
    try {
      const matches = await this.eloService.getMatches();
      this.dataSource = matches.sort((a, b) => {
        if (a.dateTime > b.dateTime) {
          return 1;
        }
        else if (a.dateTime === b.dateTime) {
          return 0;
        }

        return -1;
      })
    }
    finally {
      this.isLoaded = true;
    }
  }

  private parseTime(): Date {
    const dt = DateTime.fromISO(this.newTime!);
    return dt.toJSDate();
  }

  public async addMatch(): Promise<void> {
    this.isLoaded = false;

    const timeParsed = this.parseTime();
    const localDateTime = new Date(this.newDate!.getFullYear(), this.newDate!.getMonth(), this.newDate!.getDate(), timeParsed.getHours(), timeParsed.getMinutes(), 0, 0);
    const utcDateTime = new Date(Date.UTC(localDateTime.getUTCFullYear(), localDateTime.getUTCMonth(),
      localDateTime.getUTCDate(), localDateTime.getUTCHours(),
      localDateTime.getUTCMinutes(), localDateTime.getUTCSeconds()));

    const newMatch: Match = {
      dateTime: utcDateTime,
      leftPlayer: this.newPlayer1.playerName!,
      leftPlayerScore: this.newPlayer1.score!,
      rightPlayer: this.newPlayer2.playerName!,
      rightPlayerScore: this.newPlayer2.score!,
    };

    await this.eloService.addMatch(newMatch);
    await this.refresh();

    this.newDate = undefined;
    this.newTime = undefined;
    this.newPlayer1 = new PlayerScoreViewModel();
    this.newPlayer2 = new PlayerScoreViewModel();
  }

  public canAddMatches(): boolean {
    return this.isLoaded && this.newDate != null && this.newPlayer1.isOk && this.newPlayer2.isOk && this.newPlayer1.playerName !== this.newPlayer2.playerName && this.newTime != null;
  }

  public async deleteMatch(match: Match): Promise<void> {
    if (await this.messageBoxService.showYesNoQuestion("Question", "Do you want to delete this match?")) {
      this.isLoaded = false;
      await this.eloService.deleteMatch(match.id!);
      await this.refresh();
    }
  }
}

class PlayerScoreViewModel {
  public playerName?: string;
  public score?: number;

  public get isOk(): boolean {
    return this.playerName != null && this.score != null;
  }
}