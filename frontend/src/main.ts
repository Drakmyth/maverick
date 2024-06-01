import "./style.css";

import {
    GetHomePage,
    GetEnginesPage,
    GetIWADsPage,
    GetPageTitle,
    GetPageDescription,
    GetAddIWADModal,
    GetIWADOptionsModal,
    SelectIWADFile,
    SaveIWAD,
} from "../wailsjs/go/main/App";

declare global {
    interface Window {
        openAddIWADModal: () => void;
        closeAddIWADModal: () => void;
        openIWADOptionsModal: (event: MouseEvent) => void;
        closeIWADOptionsModal: () => void;
        navigateTo: (page: string) => void;
        mavInit: () => void;
        selectIWADFile: () => void;
        validateAddIWADForm: () => void;
        submitAddIWADForm: (event: SubmitEvent) => void;
    }
}

window.openAddIWADModal = async function () {
    let template = document.createElement("template");
    template.innerHTML = await GetAddIWADModal();
    let dialog = template.content.children[0] as HTMLDialogElement;

    let app = document.getElementById("app") as HTMLDivElement;
    app.append(dialog);

    dialog.showModal();
};

window.closeAddIWADModal = function () {
    let dialog = document.getElementById("add-iwad-dialog") as HTMLDialogElement;
    dialog.close();
    dialog.remove();
};

window.openIWADOptionsModal = async function (event: MouseEvent) {
    let template = document.createElement("template");
    template.innerHTML = await GetIWADOptionsModal();
    let dialog = template.content.children[0] as HTMLDialogElement;
    dialog.onmousedown = getDialogCoverClickHandler(dialog, closeIWADOptionsModal);

    let app = document.getElementById("app") as HTMLDivElement;
    app.append(dialog);

    let position = getIdealDialogPosition(event.target as HTMLElement, dialog);
    dialog.style.marginLeft = `${position.x}px`;
    dialog.style.marginTop = `${position.y}px`;

    window.addEventListener("resize", function resized() {
        closeIWADOptionsModal();
        window.removeEventListener("resize", resized);
    });

    dialog.showModal();
};

function getIdealDialogPosition(trigger: HTMLElement, dialog: HTMLDialogElement) {
    let triggerBounds = trigger.getBoundingClientRect();
    let dialogBounds = dialog.getBoundingClientRect();

    let viewportBounds = document.documentElement.getBoundingClientRect();

    let left = triggerBounds.right - dialogBounds.width;
    if (left < 0) {
        left = triggerBounds.left;
    }

    let top = triggerBounds.bottom;
    if (top + dialogBounds.height > viewportBounds.bottom) {
        top = triggerBounds.top - dialogBounds.height;
    }

    return new DOMPoint(left, top);
}

function getDialogCoverClickHandler(dialog: HTMLDialogElement, callback: Function) {
    return (event: MouseEvent) => {
        let rect = dialog.getBoundingClientRect();
        let clickedInDialog =
            event.clientX >= rect.left &&
            event.clientX <= rect.right &&
            event.clientY >= rect.top &&
            event.clientY <= rect.bottom;

        if (!clickedInDialog) {
            callback();
        }
    };
}

function closeIWADOptionsModal() {
    let dialog = document.getElementById("iwad-options-modal") as HTMLDialogElement;
    dialog.close();
    dialog.remove();
}

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
