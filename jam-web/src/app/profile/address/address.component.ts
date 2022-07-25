import { Component, ViewChild, OnInit, ElementRef, NgZone } from '@angular/core';
import { Geolocation } from '@ionic-native/geolocation/ngx';
import { NativeGeocoder, NativeGeocoderOptions } from '@ionic-native/native-geocoder/ngx';

declare var google;

@Component({
  selector: 'app-address',
  templateUrl: './address.component.html',
  styleUrls: ['./address.component.scss'],
})
export class AddressComponent implements OnInit {
    // @ViewChild('map',  {static: false}) mapElement: ElementRef;
    // map: any;
    address:string;
    lat: string;
    long: string;  
    autocomplete: { input: string; };
    autocompleteItems: any[];
    location: any;
    placeid: any;
    googleAutocomplete: any;
  
    constructor(
      private geolocation: Geolocation,
      private nativeGeocoder: NativeGeocoder,    
      public zone: NgZone,
    ) {
      this.googleAutocomplete = new google.maps.places.AutocompleteService();
      this.autocomplete = { input: '' };
      this.autocompleteItems = [];
    }

    ngOnInit() {
    }
  
    public async getCoordinates(): Promise<any> {
      let resp = await this.geolocation.getCurrentPosition();
      let latLng = new google.maps.LatLng(resp.coords.latitude, resp.coords.longitude);
      let mapOptions = {
        center: latLng,
        zoom: 15,
        mapTypeId: google.maps.MapTypeId.ROADMAP
      } 
      
      await this.getAddressFromCoords(resp.coords.latitude, resp.coords.longitude); 
      // this.map = new google.maps.Map(this.mapElement.nativeElement, mapOptions); 
      // this.map.addListener('tilesloaded', () => {
      //   console.log('accuracy',this.map, this.map.center.lat());
      //   this.getAddressFromCoords(this.map.center.lat(), this.map.center.lng());
      //   this.lat = this.map.center.lat()
      //   this.long = this.map.center.lng()
      // });
    }
  
    async getAddressFromCoords(latitude, longitude): Promise<any> {
      console.log("getAddressFromCoords " + latitude + " " + longitude);
      let options: NativeGeocoderOptions = {
        useLocale: true,
        maxResults: 5    
      };

      let result = await this.nativeGeocoder.reverseGeocode(latitude, longitude, options);

      this.address = "";
      let responseAddress = [];
      for (let [_, value] of Object.entries(result[0])) {
        if (value.length > 0) responseAddress.push(value); 
      }
      responseAddress.reverse();
      for (let value of responseAddress) {
        this.address += value + ", ";
      }
      this.address = this.address.slice(0, -2);
    }
    
    //AUTOCOMPLETE, SIMPLY LOAD THE PLACE USING GOOGLE PREDICTIONS AND RETURNING THE ARRAY.
    updateSearchResults(){
      if (this.autocomplete.input == '') {
        this.autocompleteItems = [];
        return;
      }
      this.googleAutocomplete.getPlacePredictions({ input: this.autocomplete.input },
      (predictions, _status) => {
        this.autocompleteItems = [];
        this.zone.run(() => {
          predictions.forEach((prediction) => {
            this.autocompleteItems.push(prediction);
          });
        });
      });
    }
    
    //wE CALL THIS FROM EACH ITEM.
    selectSearchResult(item) {    
      this.placeid = item.place_id;
      this.location = item.description;
      alert(this.location);
    }
    
    clearAutocomplete(){
      this.autocompleteItems = []
      this.autocomplete.input = ''
    }
}
