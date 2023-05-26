import { ChangeDetectorRef, Component, OnDestroy } from '@angular/core';
import { MediaMatcher } from '@angular/cdk/layout';
import { MenuItems } from '../../../shared/menu-items/menu-items';
@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: []
})
export class AppSidebarComponent implements OnDestroy {
  username:string="";
  logged:boolean=false;
  mobileQuery: MediaQueryList;

  private _mobileQueryListener: () => void;

  constructor(
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher,
    public menuItems: MenuItems
  ) {
    this.mobileQuery = media.matchMedia('(min-width: 768px)');
    this._mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this._mobileQueryListener);
  }
  ngOnInit() {
    const token = localStorage.getItem('token');
    const username = localStorage.getItem('username');
    if (token) {
      this.logged = true;
     
    }
    if(username){
      this.username=username;
    }
  }
  login() {
      this.logged = true;
    }
  logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    this.logged = false;
  }
  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this._mobileQueryListener);
  }
}
