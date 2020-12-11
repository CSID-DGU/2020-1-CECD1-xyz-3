import {
  ChangeDetectorRef,
  Component,
  OnInit,
} from '@angular/core';
import { ApiService } from 'app/core';
import * as RxOp from 'rxjs/operators';
import { TranslateService } from 'app/core/translate.service';

enum PageMode {
  VIEW,
  EDIT,
  CREATE,
}

@Component({
  selector: 'dg-position-management',
  templateUrl: './position-management.component.html',
  styleUrls: [ './position-management.component.scss' ],
})
export class PositionManagementComponent implements OnInit {
  public readonly PageMode: typeof PageMode = PageMode;
  public activePageMode: PageMode = PageMode.VIEW;

  public positions: JobPosition[] = [];
  public isFetching: boolean = false;
  public selectedPositions: JobPosition[] = [];

  public editingPoisition: JobPosition;

  public constructor(
    private readonly apiService: ApiService,
    private readonly changeDetectorRef: ChangeDetectorRef,
    private readonly translateService: TranslateService,
  ) {
    return;
  }

  public ngOnInit(): void {
    this.fetchPositions();
  }

  public onCreatePositionClick(): void {
    this.editingPoisition = {
      id: this.getUniqueId(),
      category: '',
      name: '',
      summary: '',
      descriptions: [
        { title: '', texts: [ '' ] },
      ],
    };
    this.activePageMode = PageMode.CREATE;
  }

  public onViewDetailClick(position: JobPosition): void {
    this.editingPoisition = JSON.parse(JSON.stringify(position));
    this.activePageMode = PageMode.EDIT;
  }

  public onCancelClick(): void {
    this.activePageMode = PageMode.VIEW;
  }

  public onSaveClick(): void {
    this.apiService.updatePositions([ this.editingPoisition ])
      .subscribe(() => {
        this.fetchPositions();
        this.activePageMode = PageMode.VIEW;
      }, err => {
        alert(this.translateService.translate('ERROR_SAVE_POSITION_CHANGE_FAIELD'));
      });
  }

  public onCreateClick(): void {
    this.apiService.createPositions([ this.editingPoisition ])
      .subscribe(() => {
        this.fetchPositions();
        this.activePageMode = PageMode.VIEW;
      }, err => {
        alert(this.translateService.translate('ERROR_CREATE_NEW_POSITION_FAILED'));
      });
  }

  public isSelected(position: JobPosition): boolean {
    return this.selectedPositions.indexOf(position) > -1;
  }

  public onCheckboxClick(position: JobPosition): void {
    const index = this.selectedPositions.indexOf(position);
    if (index > -1) {
      this.selectedPositions.splice(index, 1);
    } else {
      this.selectedPositions.push(position);
    }
  }

  public onDeleteButtonClick(): void {
    if (this.selectedPositions.length === 0) {
      alert(this.translateService.translate('SELECT_POSITION_TO_REMOVE'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_REMOVE_POSITION_MESSAGE').replace('$1', `${this.selectedPositions.length}`))) {
      this.apiService.removePositions(this.selectedPositions)
        .subscribe(() => {
          this.fetchPositions();
        }, err => {
          alert(this.translateService.translate('ERROR_REMOVE_POSITION_FAILED'));
        });
    }
  }

  private fetchPositions(): void {
    this.isFetching = true;
    this.apiService.getPositions()
      .pipe(
        RxOp.finalize(() => this.isFetching = false),
      )
      .subscribe(positions => {
        this.selectedPositions.splice(0);
        this.positions = positions;
        this.changeDetectorRef.markForCheck();
      }, err => {
        alert(this.translateService.translate('ERROR_FETCH_POSITION_LIST'));
      });
  }

  private getUniqueId(): number {
    return (Math.max(...[ 0, ...this.positions.map(p => p.id) ]) + 1);
  }
}
