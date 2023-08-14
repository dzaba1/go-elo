import { Injectable } from '@angular/core';
import { Match } from 'src/model/match';
import { PlayerRating } from 'src/model/playerRating';

@Injectable({
  providedIn: 'root'
})
export class EloService {

  constructor() { }

  public async getLeagueName(): Promise<string> {
    return Promise.resolve("TODO")
  }

  public async getRating(): Promise<PlayerRating[]> {
    const temp1: PlayerRating = {
      player: "player1",
      rating: 1010
    }
    const temp2: PlayerRating = {
      player: "player3",
      rating: 990
    }
    
    return Promise.resolve([temp1, temp2]);
  }

  public async addMatch(match: Match): Promise<Match> {
    return Promise.resolve(match);
  }
}
