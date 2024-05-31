import "./style.css";

import {
    GetHomePage,
    GetEnginesPage,
    GetIWADsPage,
    GetPageTitle,
    GetPageDescription,
    GetAddIWADModal,
} from "../wailsjs/go/main/App";

declare global {
    interface Window {
        openAddIWADModal: () => void;
        closeAddIWADModal: () => void;
        navigateTo: (page: string) => void;
        mavInit: () => void;
    }
}

window.openAddIWADModal = async function () {
    let openIWADModalCover = document.getElementById(
        "modal-cover"
    ) as HTMLDivElement;
    let modal = document.getElementById("modal") as HTMLDivElement;

    openIWADModalCover.classList.add("show-modal");
    modal.innerHTML = await GetAddIWADModal()
};

window.closeAddIWADModal = function () {
    let openIWADModalCover = document.getElementById(
        "modal-cover"
    ) as HTMLDivElement;
    openIWADModalCover.classList.remove("show-modal");
};

window.navigateTo = async function (page: string) {
    let appDiv = document.getElementById("page-content") as HTMLDivElement;
    let navButtons = (
        document.getElementById("navbar") as HTMLDivElement
    ).getElementsByTagName("button");
    let pageTitleHeading = document.getElementById(
        "page-title"
    ) as HTMLHeadingElement;
    let pageDescriptionParagraph = document.getElementById(
        "page-description"
    ) as HTMLParagraphElement;

    Array.from(navButtons).forEach((btn) => {
        if (btn.id === `nav-${page}`) {
            btn.classList.add("selected");
        } else {
            btn.classList.remove("selected");
        }
    });

    pageTitleHeading.innerText = await GetPageTitle(page);
    pageDescriptionParagraph.innerText = await GetPageDescription(page);

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
