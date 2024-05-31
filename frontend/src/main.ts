import "./style.css";
import "./app.css";

import {
    GetHomePage,
    GetEnginesPage,
    GetIWADsPage,
} from "../wailsjs/go/main/App";

declare global {
    interface Window {
        openAddIWADModal: () => void;
        closeAddIWADModal: () => void;
        navigateTo: (page: string) => void;
        mavInit: () => void;
    }
}

window.openAddIWADModal = function () {
    let openIWADModalCover = document.getElementsByClassName(
        "modal-cover"
    )[0] as HTMLDivElement;
    openIWADModalCover.classList.add("show-modal");
};

window.closeAddIWADModal = function () {
    let openIWADModalCover = document.getElementsByClassName(
        "modal-cover"
    )[0] as HTMLDivElement;
    openIWADModalCover.classList.remove("show-modal");
};

window.navigateTo = async function (page: string) {
    let appDiv = document.getElementById("app") as HTMLDivElement;
    switch (page) {
        case "home":
            appDiv.innerHTML = await GetHomePage();
            break;
        case "engines":
            appDiv.innerHTML = await GetEnginesPage();
            break;
        case "iwads":
            appDiv.innerHTML = await GetIWADsPage();
            break;
    }
};

window.mavInit = function () {
    window.navigateTo("home");
};
