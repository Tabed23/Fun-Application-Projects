import { BookService } from './book.service';
import { Args, Int, Mutation, Query, Resolver } from "@nestjs/graphql";
import { Book } from "./book.schema";
import {Book as BookModel} from "../graphql";
import { BookArgs } from './args/book.args';


@Resolver((of) => Book)
export class BookResolver{


    constructor(private readonly bookService: BookService){}

    @Query(returns => [Book])
    public getAllBooks(): BookModel[] {
       return this.bookService.findAllBooks();
    }


    @Query(returns => Book, {nullable: true})
    public getBookByID(@Args({ name: "bookID", type: () => Int}) id: number): BookModel {
        return this.bookService.findBookById(id);
    }


    @Mutation(returns => String)
    public deleteBookByID(@Args({ name: "bookID", type: () => Int}) id: number): String {
        return this.bookService.deleteBook(id);
        }


    @Mutation(returns => String)
    public createBook(@Args("addBookArgs") addBookArgs: BookArgs): String {
        return this.bookService.addBook(addBookArgs)
    }


    @Mutation(returns => String)
    public UpdateBook(@Args({ name: "bookID", type: () => Int}) id: number, @Args("updateBookArgs") updateBookArgs: BookArgs): String {
        return this.bookService.updateBook(id, updateBookArgs)
    }
}
