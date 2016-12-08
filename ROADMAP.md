Ana Roadmap
===========

This is a general draft document for thoughts and todo's, without any structure to it.

### What's cooking?

- Hand out unique ID to each visitor
- Reference site URL when tracking.
- Reference path & title when tracking (indexed by path, update title when changes)
- Track referrals, use tables from aforementioned points.
- CLI commands for CRUD user.
- Bulk process tracking requests (Redis or in-memory?)
- Allow sorting in table overviews.
- Choose a OS license & settle on name.
- Envelope API responses & perhaps return total in table overview?
- Track canonical URL's.
- Show referrals.
- Geolocate unknown IP addresses periodically.
- Mask last part of IP address.

### Key metrics

- Unique visits per day (in period)
- Pageviews per day (in period)
- Demographic
  - Country
  - Browser + version
  - Screen resolutions
- Acquisition
  - Referral's
  - Search keywords

```
  // stmt2, _ := db.Conn.Prepare("INSERT INTO users(email, password) VALUES(?, ?)")
  // hash, _ := bcrypt.GenerateFromPassword([]byte(l.Password), 10)
  // stmt2.Exec(l.Email, hash)
```
