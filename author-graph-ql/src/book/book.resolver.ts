import { BookService } from './book.service';
import { Args, Int, Mutation, Query, Resolver } from "@nestjs/graphql";
import { Book } from "./book.schema";
import {  CreateBookArgs, UpdateBookArgs } from './args/book.args';


@Resolver((of) => Book)
export class BookResolver{


    constructor( private readonly bookService: BookService){}

    @Query(returns => [Book])
    getAllBooks() {
       return this.bookService.findAllBooks();
    }


    @Query(returns => Book, {nullable: true})
    getBookByID(@Args({ name: "bookID", type: () => Int}) id: number) {
        return this.bookService.findBookById(id)
    }


    @Mutation(returns => String)
    deleteBookByID(@Args({ name: "bookID", type: () => Int}) id: number) {
        return this.bookService.deleteBook(id);
    }


    @Mutation(returns => String)
    createBook(@Args("addBookArgs") addBookArgs: CreateBookArgs) {
        return this.bookService.addBook(addBookArgs)
    }


    @Mutation(returns => String)
    UpdateBook(@Args("updateBookArgs") updateBookArgs: UpdateBookArgs) {
        return this.bookService.updateBook(updateBookArgs)
    }
}
