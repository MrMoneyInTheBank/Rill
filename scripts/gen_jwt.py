from pathlib import Path


from typing import Optional
import typing

from coinbase import jwt_generator
import dotenv

REQUEST_METHOD: typing.Final[str] = "GET"
REQUEST_PATH: typing.Final[str] = "/api/v3/brokerage/products"


def get_project_root() -> Path:
    current_path: Path = Path(__file__).resolve()

    for parent in current_path.parents:
        if (parent / ".git").exists():
            return parent

    return current_path.parent


def main() -> None:

    try:
        dotenv_path: Path = Path(dotenv.find_dotenv(raise_error_if_not_found=True))
    except FileNotFoundError:
        raise FileNotFoundError("Missing .env file")

    api_key: Optional[str] = dotenv.get_key(dotenv_path, "API_KEY")
    api_secret: Optional[str] = dotenv.get_key(dotenv_path, "API_SECRET")

    if not api_key:
        raise ValueError("Missing required environment variable: API_KEY")
    if not api_secret:
        raise ValueError("Missing required environment variable: API_SECRET")

    try:
        jwt_uri: str = jwt_generator.format_jwt_uri(REQUEST_METHOD, REQUEST_PATH)
        rest_jwt_token: str = jwt_generator.build_rest_jwt(jwt_uri, api_key, api_secret)
        ws_jwt_token: str = jwt_generator.build_ws_jwt(api_key, api_secret)

        dotenv.set_key(dotenv_path, "REST_JWT", rest_jwt_token)
        dotenv.set_key(dotenv_path, "WS_JWT", ws_jwt_token)
    except Exception as e:
        raise RuntimeError(
            f"Failed to generate Coinbase JWT. Verify API keys are correct.\nOriginal Error: {e}"
        ) from e

    print("Added keys to env file.")


if __name__ == "__main__":
    main()
