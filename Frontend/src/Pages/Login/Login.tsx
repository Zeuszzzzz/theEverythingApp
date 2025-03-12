export default function Login(){
    return (
        <>
            <form method="post">
                <input type="email" name="EMAIL" placeholder="Email" id="" />
                <input type="password" name="PASSWORD" placeholder="Password" id="" />
                <input type="submit" value="Submit" />
            </form>
        </>
    );
}