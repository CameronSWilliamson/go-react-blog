export const logIn = (username: string) => {
    localStorage.setItem("username", username);
    localStorage.setItem("loggedIn", "true");
};

export const logOut = () => {
    localStorage.removeItem("username");
    localStorage.removeItem("loggedIn");
};

interface IAuth {
    username: string;
    loggedIn: boolean;
}

export const checkLogin = (): IAuth => {
    const username = localStorage.getItem("username");
    const loggedIn = localStorage.getItem("loggedIn");

    if (username && loggedIn) {
        return {
            username,
            loggedIn: true,
        } as IAuth;
    } else {
        return {
            username: "",
            loggedIn: false,
        } as IAuth;
    }
};
