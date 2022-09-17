export class Profile {
    id: string
    name: string
    email: string
    gender: string
    instruments: string[]

    constructor() {
        this.instruments = [];
    }

    addInstrument(instrument: string) {
        this.instruments.push(instrument);
    }
}