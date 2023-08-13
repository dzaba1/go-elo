import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { PlayerRating } from 'src/model/playerRating';
import { EloService } from 'src/services/elo.service';

@Component({
  selector: 'app-rating',
  templateUrl: './rating.component.html',
  styleUrls: ['./rating.component.sass']
})
export class RatingComponent implements OnInit {
  public dataSource: PlayerRating[] = [];
  public displayedColumns = ["player", "rating"];
  public isLoaded = false;

  constructor(private eloService: EloService) {

  }

  public async ngOnInit(): Promise<void> {
    this.dataSource = await this.eloService.getRating();
    this.isLoaded = true;
  }
}
