import './style.css';
import './app.css';

// import {Greet} from '../wailsjs/go/main/App';

declare global {
    interface Window {
        openAddIWADModal: () => void;
        closeAddIWADModal: () => void;
    }
}

window.openAddIWADModal = function () {
    let openIWADModalCover = (document.getElementsByClassName("modal-cover")[0] as HTMLDivElement);
    openIWADModalCover.classList.add("show-modal");
}

window.closeAddIWADModal = function () {
    let openIWADModalCover = (document.getElementsByClassName("modal-cover")[0] as HTMLDivElement);
    openIWADModalCover.classList.remove("show-modal");
}
