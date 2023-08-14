import { Component, OnInit } from '@angular/core';
import { Match } from 'src/model/match';
import { EloService } from 'src/services/elo.service';

@Component({
  selector: 'app-matches',
  templateUrl: './matches.component.html',
  styleUrls: ['./matches.component.sass']
})
export class MatchesComponent implements OnInit {
  public displayedColumns = ["dateTime", "player1", "score", "player2"];
  public dataSource: Match[] = [];
  public isLoaded = false;
  public newDate?: Date;
  public newPlayer1 = new PlayerScoreViewModel();
  public newPlayer2 = new PlayerScoreViewModel();

  constructor(private eloService: EloService) {

  }

  public async ngOnInit(): Promise<void> {
    await this.refresh();
  }

  private async refresh(): Promise<void> {
    this.isLoaded = false;
    try {
      this.dataSource = await this.eloService.getMatches();
    }
    finally {
      this.isLoaded = true;
    }
  }

  public async addMatch(): Promise<void> {
    this.isLoaded = false;
    
    const newMatch: Match = {
      dateTime: this.newDate!,
      leftPlayer: this.newPlayer1.playerName!,
      leftPlayerScore: this.newPlayer1.score!,
      rightPlayer: this.newPlayer2.playerName!,
      rightPlayerScore: this.newPlayer2.score!,
    };

    await this.eloService.addMatch(newMatch);
    await this.refresh();

    this.newDate = undefined;
    this.newPlayer1 = new PlayerScoreViewModel();
    this.newPlayer2 = new PlayerScoreViewModel();
  }

  public canAddMatches(): boolean {
    return this.isLoaded && this.newDate != null && this.newPlayer1.isOk && this.newPlayer2.isOk && this.newPlayer1.playerName !== this.newPlayer2.playerName;
  }
}

class PlayerScoreViewModel {
  public playerName?: string;
  public score?: number;

  public get isOk(): boolean {
    return this.playerName != null && this.score != null;
  }
}