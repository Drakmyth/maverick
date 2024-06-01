import "./style.css";

import {
    GetHomePage,
    GetEnginesPage,
    GetIWADsPage,
    GetPageTitle,
    GetPageDescription,
    GetAddIWADModal,
    SelectIWADFile,
    SaveIWAD,
} from "../wailsjs/go/main/App";

declare global {
    interface Window {
        openAddIWADModal: () => void;
        closeAddIWADModal: () => void;
        navigateTo: (page: string) => void;
        mavInit: () => void;
        selectIWADFile: () => void;
        validateAddIWADForm: () => void;
        submitAddIWADForm: (event: SubmitEvent) => void;
    }
}

window.openAddIWADModal = async function () {
    let dialog = document.getElementById("dialog") as HTMLDialogElement;
    dialog.innerHTML = await GetAddIWADModal();
    dialog.showModal();
};

window.closeAddIWADModal = function () {
    let dialog = document.getElementById("dialog") as HTMLDialogElement;
    dialog.close();
    dialog.innerHTML = "";
};

window.navigateTo = async function (page: string) {
    let appDiv = document.getElementById("page-content") as HTMLDivElement;
    let navButtons = (document.getElementById("navbar") as HTMLDivElement).getElementsByTagName("button");
    let pageTitleHeading = document.getElementById("page-title") as HTMLHeadingElement;
    let pageDescriptionParagraph = document.getElementById("page-description") as HTMLParagraphElement;

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

window.selectIWADFile = async function () {
    let pathInput = document.getElementById("iwad-file-txt") as HTMLInputElement;
    pathInput.value = await SelectIWADFile();
    window.validateAddIWADForm();
};

window.validateAddIWADForm = function () {
    let nameInput = document.getElementById("iwad-name-txt") as HTMLInputElement;
    let pathInput = document.getElementById("iwad-file-txt") as HTMLInputElement;
    let submitButton = document.getElementById("iwad-submit") as HTMLButtonElement;

    let valid = Boolean(nameInput.value.trim()) && Boolean(pathInput.value.trim());
    submitButton.disabled = !valid;
};

window.submitAddIWADForm = async function (event: SubmitEvent) {
    event.preventDefault();

    let nameInput = document.getElementById("iwad-name-txt") as HTMLInputElement;
    let pathInput = document.getElementById("iwad-file-txt") as HTMLInputElement;

    await SaveIWAD({
        Name: nameInput.value,
        Path: pathInput.value,
    });

    window.closeAddIWADModal();
    window.navigateTo("iwads");
};
