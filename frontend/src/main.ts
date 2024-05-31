import './style.css';
import './app.css';

import {GetHomePage, GetEnginesPage, GetIWADsPage} from '../wailsjs/go/main/App';

declare global {
    interface Window {
        openAddIWADModal: () => void;
        closeAddIWADModal: () => void;
        navigateToHome: () => void;
        navigateToEngines: () => void;
        navigateToIWADs: () => void;
        mavInit: () => void;
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

window.navigateToHome = async function () {
    let appDiv = document.getElementById("app") as HTMLDivElement;
    appDiv.innerHTML = await GetHomePage()
}

window.navigateToEngines = async function () {
    let appDiv = document.getElementById("app") as HTMLDivElement;
    appDiv.innerHTML = await GetEnginesPage()
}

window.navigateToIWADs = async function () {
    let appDiv = document.getElementById("app") as HTMLDivElement;
    appDiv.innerHTML = await GetIWADsPage()
}

window.mavInit = function () {
    window.navigateToHome()
}