import "./app.css";
import "./style.css";
import "./modals/modal";

import {
    GetContent,
    GetPageTitle,
    GetPageDescription,
    SelectIWADFile,
    MoveIWADUp,
    MoveIWADDown,
} from "../wailsjs/go/main/App";
import { AddIWADModal } from "./modals/add-iwad-modal";
import { ModifyIWADModal } from "./modals/modify-iwad-modal";
import { RemoveIWADModal } from "./modals/remove-iwad-modal";
import { IWADOptionsModal } from "./modals/iwad-options-modal";
import { AddEngineModal } from "./modals/add-engine-modal";

declare global {
    interface Window {
        addIWADModal: AddIWADModal;
        modifyIWADModal: ModifyIWADModal;
        removeIWADModal: RemoveIWADModal;
        iwadOptionsModal: IWADOptionsModal;
        addEngineModal: AddEngineModal;
        navigateTo(page: string): void;
        mavInit(): void;
        selectIWADFile(): void;
        validateAddIWADForm(): void;
        validateAddEngineForm(): void;
        moveIWADUp(iwadId: string): void;
        moveIWADDown(iwadId: string): void;
    }
}

window.addIWADModal = new AddIWADModal();
window.modifyIWADModal = new ModifyIWADModal();
window.removeIWADModal = new RemoveIWADModal();
window.iwadOptionsModal = new IWADOptionsModal();
window.addEngineModal = new AddEngineModal();

window.navigateTo = async function (page: string) {
    let appDiv = document.getElementById("page-content") as HTMLDivElement;
    let navButtons = (document.getElementById("navbar") as HTMLDivElement).getElementsByTagName("button");
    let pageTitleHeading = document.getElementById("page-title") as HTMLHeadingElement;
    let pageDescriptionParagraph = document.getElementById("page-description") as HTMLParagraphElement;

    Array.from(navButtons).forEach((btn) => {
        if (btn.id === `nav-${page}`) {
            btn.classList.add("btn-selected");
        } else {
            btn.classList.remove("btn-selected");
        }
    });

    pageTitleHeading.innerText = await GetPageTitle(page);
    pageDescriptionParagraph.innerText = await GetPageDescription(page);

    switch (page) {
        case "home":
            appDiv.innerHTML = await GetContent("home-page");
            break;
        case "engines":
            appDiv.innerHTML = await GetContent("engines-page");
            break;
        case "iwads":
            appDiv.innerHTML = await GetContent("iwads-page");
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

window.validateAddEngineForm = function () {
    let nameInput = document.getElementById("engine-name-txt") as HTMLInputElement;
    let pathInput = document.getElementById("engine-file-txt") as HTMLInputElement;
    let submitButton = document.getElementById("engine-submit") as HTMLButtonElement;

    let valid = Boolean(nameInput.value.trim()) && Boolean(pathInput.value.trim());
    submitButton.disabled = !valid;
};

window.moveIWADUp = async function (iwadId: string) {
    let success = await MoveIWADUp(iwadId);
    if (!success) {
        return;
    }

    window.iwadOptionsModal.close();
    window.navigateTo("iwads");
};
window.moveIWADDown = async function (iwadId: string) {
    let success = await MoveIWADDown(iwadId);
    if (!success) {
        return;
    }

    window.iwadOptionsModal.close();
    window.navigateTo("iwads");
};
