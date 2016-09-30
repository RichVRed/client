// Copyright 2016 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package pvl

var hardcodedPVLString = `
{
  "pvl_version": 1,
  "revision": 1,
  "services": {
    "coinbase": [
      [
        { "fill": { "into": "our_url", "with": "https://coinbase.com/%{username_service}/public-key" } },
        { "fetch": { "from": "our_url", "kind": "html" } },
        {
          "selector_css": {
            "error": ["FAILED_PARSE", "Couldn't find a div $(pre.statement)"],
            "into": "haystack",
            "selectors": ["pre.statement", 0]
          }
        },
        {
          "assert_find_base64": { "haystack": "haystack", "needle": "sig" },
          "error": ["TEXT_NOT_FOUND", "signature not found in body"]
        }
      ]
    ],
    "dns": [
      [
        {
          "assert_regex_match": {
            "error": ["CONTENT_FAILURE", "matching DNS entry not found"],
            "from": "txt",
            "pattern": "^keybase-site-verification=%{sig_id_medium}$"
          }
        }
      ]
    ],
    "facebook": [
      [
        {
          "regex_capture": {
            "error": [
              "BAD_API_URL",
              "Bad hint from server; URL should start with 'https://m.facebook.com/%{username_service}/posts/', received '%{hint_url}'"
            ],
            "from": "hint_url",
            "into": ["username_from_url"],
            "pattern": "^https://m\\.facebook\\.com/([^/]*)/posts/.*$"
          }
        },
        {
          "assert_compare": {
            "a": "username_from_url",
            "b": "username_service",
            "cmp": "stripdots-then-cicmp",
            "error": [
              "BAD_API_URL",
              "Bad hint from server; username in URL match '%{username_service}', received '%{username_from_url}'"
            ]
          }
        },
        { "fetch": { "from": "hint_url", "kind": "html" } },
        {
          "selector_css": {
            "attr": "href",
            "error": ["FAILED_PARSE", "Couldn't find username href"],
            "into": "username_link",
            "selectors": ["#m_story_permalink_view > div:first-child > div:first-child > div:first-child h3", 0, "a", 0]
          }
        },
        {
          "parse_url": {
            "error": ["FAILED_PARSE", "Failed to parse username URL: %{username_link}"],
            "from": "username_link",
            "path": "path"
          }
        },
        {
          "regex_capture": {
            "error": ["FAILED_PARSE", "Username URL has no path"],
            "from": "path",
            "into": ["split_path_1"],
            "pattern": "^[^/]*/([^/]*)$"
          }
        },
        {
          "assert_compare": {
            "a": "split_path_1",
            "b": "username_service",
            "cmp": "stripdots-then-cicmp",
            "error": ["BAD_USERNAME", "Usernames don't match '%{split_path_1}' vs '%{username_service}'"]
          }
        },
        {
          "selector_css": {
            "error": ["FAILED_PARSE", "Couldn't find proof text header"],
            "into": "header",
            "selectors": ["#m_story_permalink_view > div:first-child > div:first-child > div:first-child h3", 1]
          }
        },
        { "whitespace_normalize": { "from": "header", "into": "header_nw" } },
        {
          "assert_regex_match": {
            "error": ["TEXT_NOT_FOUND", "Proof text not found: '' != ''"],
            "from": "header_nw",
            "pattern": "^Verifying myself: I am %{username_keybase} on Keybase\\.io\\. %{sig_id_medium}$"
          }
        }
      ]
    ],
    "generic_web_site": [
      [
        {
          "assert_regex_match": {
            "error": ["BAD_API_URL", "Bad hint from server; didn't recognize API url: \"%{active_string}\""],
            "from": "hint_url",
            "pattern": "^%{protocol}://%{hostname}/(?:\\.well-known/keybase\\.txt|keybase\\.txt)$"
          }
        },
        { "fetch": { "from": "hint_url", "into": "blob", "kind": "string" } },
        {
          "assert_find_base64": { "error": ["TEXT_NOT_FOUND", "signature not found in body"], "haystack": "blob", "needle": "sig" }
        }
      ]
    ],
    "github": [
      [
        {
          "regex_capture": {
            "error": [
              "BAD_API_URL",
              "Bad hint from server; URL should start with either https://gist.github.com OR https://gist.githubusercontent.com"
            ],
            "from": "hint_url",
            "into": ["username_from_url"],
            "pattern": "^https://gist\\.github(?:usercontent)?\\.com/([^/]*)/.*$"
          }
        },
        {
          "assert_compare": {
            "a": "username_from_url",
            "b": "username_service",
            "cmp": "cicmp",
            "error": [
              "BAD_API_URL",
              "Bad hint from server; URL should contain username matching %{username_service}; got %{username_from_url}"
            ]
          }
        },
        { "fetch": { "from": "hint_url", "into": "haystack", "kind": "string" } },
        {
          "assert_find_base64": { "haystack": "haystack", "needle": "sig" },
          "error": ["TEXT_NOT_FOUND", "signature not found in body"]
        }
      ]
    ],
    "hackernews": [
      [
        {
          "regex_capture": {
            "error": [
              "BAD_API_URL",
              "Bad hint from server; URL should match https://hacker-news.firebaseio.com/v0/user/%{username_service}/about.json"
            ],
            "from": "hint_url",
            "into": ["username_from_url"],
            "pattern": "^https://hacker-news\\.firebaseio\\.com/v0/user/([^/]+)/about.json$"
          }
        },
        {
          "assert_compare": {
            "a": "username_from_url",
            "b": "username_service",
            "cmp": "cicmp",
            "error": [
              "BAD_API_URL",
              "Bad hint from server; URL should contain username matching %{username_service}; got %{username_from_url}"
            ]
          }
        },
        { "fetch": { "from": "hint_url", "into": "profile", "kind": "string" } },
        {
          "assert_regex_match": {
            "error": ["TEXT_NOT_FOUND", "Posted text does not include signature '%{sig_id_medium}'"],
            "from": "profile",
            "pattern": "^.*%{sig_id_medium}.*$"
          }
        }
      ]
    ],
    "reddit": [
      [
        {
          "regex_capture": {
            "error": ["BAD_API_URL", "URL should start with 'https://www.reddit.com/r/keybaseproofs'"],
            "from": "hint_url",
            "into": ["subreddit_from_url", "path_remainder"],
            "pattern": "^https://www.reddit.com/r/([^/]+)/(.*)$"
          }
        },
        {
          "assert_regex_match": {
            "case_insensitive": true,
            "error": ["BAD_API_URL", "URL contained wrong subreddit '%{subreddit_from_url}' !+ 'keybaseproofs'"],
            "from": "subreddit_from_url",
            "pattern": "^keybaseproofs$"
          }
        },
        { "fetch": { "from": "hint_url", "kind": "json" } },
        {
          "selector_json": {
            "error": ["CONTENT_MISSING", "Could not find 'kind' in json"],
            "into": "kind",
            "selectors": [0, "kind"]
          }
        },
        {
          "assert_regex_match": {
            "error": ["CONTENT_FAILURE", "Wanted a post of type 'Listing', but got %{kind}"],
            "from": "kind",
            "pattern": "^Listing$"
          }
        },
        {
          "selector_json": {
            "error": ["CONTENT_MISSING", "Could not find inner 'kind' in json"],
            "into": "inner_kind",
            "selectors": [0, "data", "children", 0, "kind"]
          }
        },
        {
          "assert_regex_match": {
            "error": ["CONTENT_FAILURE", "Wanted a child of type 't3' but got %{inner_kind}"],
            "from": "inner_kind",
            "pattern": "^t3$"
          }
        },
        {
          "selector_json": {
            "error": ["CONTENT_MISSING", "Could not find 'subreddit' in json"],
            "into": "subreddit_from_json",
            "selectors": [0, "data", "children", 0, "data", "subreddit"]
          }
        },
        {
          "assert_regex_match": {
            "case_insensitive": true,
            "error": ["CONTENT_FAILURE", "Wrong subreddti %{subreddit_from_json}"],
            "from": "subreddit_from_json",
            "pattern": "^keybaseproofs$"
          }
        },
        {
          "selector_json": {
            "error": ["CONTENT_MISSING", "Could not find author in json"],
            "into": "author",
            "selectors": [0, "data", "children", 0, "data", "author"]
          }
        },
        {
          "assert_compare": {
            "a": "author",
            "b": "username_service",
            "cmp": "cicmp",
            "error": ["BAD_USERNAME", "Bad post author; wanted '%{username_service} but got '%{author}'"]
          }
        },
        {
          "selector_json": {
            "error": ["CONTENT_MISSING", "Could not find title in json"],
            "into": "title",
            "selectors": [0, "data", "children", 0, "data", "title"]
          }
        },
        {
          "assert_regex_match": {
            "error": ["TITLE_NOT_FOUND", "Missing signature ID (%{sig_id_medium})) in post title '%{title}'"],
            "from": "title",
            "pattern": "^.*%{sig_id_medium}.*$"
          }
        },
        {
          "selector_json": {
            "error": ["CONTENT_MISSING", "Could not find selftext in json"],
            "into": "selftext",
            "selectors": [0, "data", "children", 0, "data", "selftext"]
          }
        },
        {
          "assert_find_base64": {
            "error": ["TEXT_NOT_FOUND", "signature not found in body"],
            "haystack": "selftext",
            "needle": "sig"
          }
        }
      ]
    ],
    "rooter": [
      [
        {
          "assert_regex_match": {
            "case_insensitive": true,
            "pattern": "^https?://[\\w:_\\-\\.]+/_/api/1\\.0/rooter/%{username_service}/.*$"
          }
        },
        { "fetch": { "kind": "json" } },
        { "selector_json": { "into": "name", "selectors": ["status", "name"] } },
        { "assert_regex_match": { "case_insensitive": true, "from": "name", "pattern": "^ok$" } },
        { "selector_json": { "into": "post", "selectors": ["toot", "post"] } },
        { "assert_regex_match": { "from": "post", "pattern": "^.*%{sig_id_medium}.*$" } }
      ]
    ],
    "twitter": [
      [
        {
          "regex_capture": {
            "error": [
              "BAD_API_URL",
              "Bad hint from server; URL should start with 'https://twitter.com/%{username_service}/'"
            ],
            "from": "hint_url",
            "into": ["username_from_url"],
            "pattern": "^https://twitter\\.com/([^/]+)/.*$"
          }
        },
        {
          "assert_compare": {
            "a": "username_from_url",
            "b": "username_service",
            "cmp": "cicmp",
            "error": [
              "BAD_API_URL",
              "Bad hint from server; URL should contain username matching %{username_service}; got %{username_from_url}"
            ]
          }
        },
        { "fetch": { "from": "hint_url", "kind": "html" } },
        {
          "selector_css": {
            "attr": "data-screen-name",
            "error": ["FAILED_PARSE", "Couldn't find a div $(div.permalink-tweet-container div.permalink-tweet).eq(0)"],
            "into": "data_screen_name",
            "selectors": ["div.permalink-tweet-container div.permalink-tweet", 0]
          }
        },
        {
          "assert_compare": {
            "a": "data_screen_name",
            "b": "username_service",
            "cmp": "cicmp",
            "error": ["BAD_USERNAME", "Bad post authored: wanted ${username_service} but got %{data_screen_name}"]
          }
        },
        {
          "selector_css": {
            "error": ["CONTENT_MISSING", "Missing <div class='tweet-text'> container for tweet"],
            "into": "tweet_contents",
            "selectors": ["div.permalink-tweet-container div.permalink-tweet", 0, "p.tweet-text", 0]
          }
        },
        { "whitespace_normalize": { "from": "tweet_contents", "into": "tweet_contents_nw" } },
        {
          "regex_capture": {
            "error": [
              "DELETED",
              "Could not find 'Verifying myself: I am %{username_keybase} on Keybase.io. %{sig_id_short}'"
            ],
            "from": "tweet_contents_nw",
            "into": ["username_from_tweet_contents", "sig_from_tweet_contents"],
            "pattern": "^ *(?:@[a-zA-Z0-9_-]+\\s*)* *Verifying myself: I am ([A-Za-z0-9_]+) on Keybase\\.io\\. (\\S+) */.*$"
          }
        },
        {
          "assert_compare": {
            "a": "username_from_tweet_contents",
            "b": "username_keybase",
            "cmp": "cicmp",
            "error": [
              "BAD_USERNAME",
              "Wrong username in tweet '%{username_from_tweet_contents}' should be '%{username_keybase}'"
            ]
          }
        },
        {
          "assert_regex_match": {
            "error": ["TEXT_NOT_FOUND", "Could not find sig '%{sig_from_tweet_contents}' != '%{sig_id_short}'"],
            "from": "sig_from_tweet_contents",
            "pattern": "^%{sig_id_short}$"
          }
        }
      ]
    ]
  }
}
`

// GetHardcodedPvlString returns the unparsed pvl
func GetHardcodedPvlString() string {
	return hardcodedPVLString
}
