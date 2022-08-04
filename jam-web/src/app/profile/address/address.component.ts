import { Component, OnInit, NgZone } from '@angular/core';

declare var google;

@Component({
  selector: 'app-address',
  templateUrl: './address.component.html',
  styleUrls: ['./address.component.scss'],
})
export class AddressComponent implements OnInit {
    address:string;
    lat: string;
    long: string;  
    autocomplete: { input: string; };
    autocompleteItems: any[];
    location: any;
    placeid: any;
    googleAutocomplete: any;
  
    constructor(
      public zone: NgZone,
    ) {
      this.googleAutocomplete = new google.maps.places.AutocompleteService();
      this.autocomplete = { input: '' };
      this.autocompleteItems = [];
    }

    ngOnInit() {
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
    
    selectSearchResult(item) {
      this.placeid = item.place_id;
      this.location = item.description;

      this.autocomplete.input = this.location;
      this.autocompleteItems = [];
    }
    
    clearAutocomplete() {
      this.autocompleteItems = [];
      this.autocomplete.input = '';
    }
}
