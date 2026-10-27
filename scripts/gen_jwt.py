import os
from pathlib import Path
from typing import Optional

from coinbase import jwt_generator
from dotenv import load_dotenv


def get_project_root() -> Path:
    current_path: Path = Path(__file__).resolve()

    for parent in current_path.parents:
        if (parent / ".git").exists():
            return parent

    return current_path.parent


def get_dotenv_path() -> Path:
    return get_project_root() / ".env"


def main() -> None:

    dotenv_path: Path = get_dotenv_path()
    if not dotenv_path.exists():
        raise FileNotFoundError("Missing .env file")

    load_dotenv(dotenv_path=dotenv_path)

    api_key: Optional[str] = os.getenv("API_KEY")
    api_secret: Optional[str] = os.getenv("API_SECRET")

    if not api_key:
        raise ValueError("Missing required environment variable: API_KEY")
    if not api_secret:
        raise ValueError("Missing required environment variable: API_SECRET")

    try:
        jwt_token: str = jwt_generator.build_ws_jwt(api_key, api_secret)
    except Exception as e:
        raise RuntimeError(
            f"Failed to generate Coinbase JWT. Verify API keys are correct.\nOriginal Error: {e}"
        ) from e

    print(f"JWT_TOKEN={jwt_token}")


if __name__ == "__main__":
    main()
