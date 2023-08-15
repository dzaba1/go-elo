import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MessageBoxData } from '../app/message-box/messageBoxData';
import { MessageBoxButtons } from 'src/app/message-box/messageBoxButtons';
import { MessageBoxComponent } from 'src/app/message-box/message-box.component';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MessageBoxService {

  constructor(private dialog: MatDialog) { }

  public async showYesNoQuestion(title: string, question: string): Promise<boolean> {
    return (await this.show(title, question, MessageBoxButtons.YesNo)) === true;
  }

  public async show(title: string, content: string, buttons: MessageBoxButtons): Promise<boolean | undefined> {
    const msgData: MessageBoxData = {
      buttons: buttons,
      title: title,
      content: content
    }

    const dialogRef = this.dialog.open<MessageBoxComponent, MessageBoxData, boolean>(MessageBoxComponent, {
      data: msgData
    });
    const obs = dialogRef.afterClosed();
    return await firstValueFrom(obs);
  }
}
