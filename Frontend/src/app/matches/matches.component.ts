import { Component, OnInit } from '@angular/core';
import { Match } from 'src/model/match';
import { EloService } from 'src/services/elo.service';

@Component({
  selector: 'app-matches',
  templateUrl: './matches.component.html',
  styleUrls: ['./matches.component.sass']
})
export class MatchesComponent implements OnInit {
  public displayedColumns = [];
  public dataSource: Match[] = [];
  public isLoaded = false;

  constructor(private eloService: EloService) {

  }

  public async ngOnInit(): Promise<void> {
    await this.refresh();
  }

  private async refresh(): Promise<void> {
    this.isLoaded = false;
    try {
      
    }
    finally {
      this.isLoaded = true;
    }
  }
}
