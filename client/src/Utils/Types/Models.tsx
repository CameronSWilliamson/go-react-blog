export interface Category {
    category_name: string;
}

export interface Post {
    title: string;
    content: string;
    username: string;
    postid: number;
}

export interface User {
    bio: string;
    email: string;
    joindate: string;
    username: string;
}
