/**
 * Used to pass in username as a prop
 */
export interface UsernameProp {
    username: string;
    setUsername: React.Dispatch<React.SetStateAction<string>>;
    loggedIn: boolean;
    setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>;
}
