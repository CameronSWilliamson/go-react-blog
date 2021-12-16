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

export interface Comment {
    commentid: number;
    postid: number;
    username: string;
    content: string;
}
