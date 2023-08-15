import { Component, Inject, Input } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MessageBoxData } from './messageBoxData';
import { MessageBoxButtons } from './messageBoxButtons';

@Component({
  selector: 'app-message-box',
  templateUrl: './message-box.component.html',
  styleUrls: ['./message-box.component.sass']
})
export class MessageBoxComponent {

  public messageBoxButtonsType = MessageBoxButtons;

  constructor(
    public dialogRef: MatDialogRef<MessageBoxComponent>,
    @Inject(MAT_DIALOG_DATA) public data: MessageBoxData,
  ) {}
}
