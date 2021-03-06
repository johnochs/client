// @flow
import * as React from 'react'
import {
  Avatar,
  Icon,
  Text,
  Box,
  iconCastPlatformStyles,
  type OverlayParentProps,
} from '../../../../../common-adapters'
import {
  globalStyles,
  globalMargins,
  globalColors,
  isMobile,
  platformStyles,
  styleSheetCreate,
  collapseStyles,
} from '../../../../../styles'
import SendIndicator from '../chat-send'
import MessagePopup from '../../message-popup'
import ExplodingHeightRetainer from '../exploding-height-retainer'
import ExplodingMeta from '../exploding-meta'
import ReactButton from '../../react-button/container'

import type {WrapperAuthorProps} from '../index.types'

const colorForAuthor = (user: string, isYou: boolean, isFollowing: boolean, isBroken: boolean) => {
  if (isYou) {
    return globalColors.black_75
  }

  if (isBroken) {
    return globalColors.red
  }
  return isFollowing ? globalColors.green2 : globalColors.blue
}

const UserAvatar = ({author, onAuthorClick}) => (
  <Box style={styles.userAvatar}>
    <Avatar size={32} username={author} skipBackground={true} onClick={onAuthorClick} />
  </Box>
)

const Username = ({username, isYou, isFollowing, isBroken, onClick}) => {
  const style = collapseStyles([
    styles.username,
    isYou && styles.usernameYou,
    {color: colorForAuthor(username, isYou, isFollowing, isBroken)},
  ])
  return (
    <Text type="BodySmallSemibold" onClick={onClick} className="hover-underline" style={style}>
      {username}
    </Text>
  )
}

const MenuButtons = ({conversationIDKey, onClick, ordinal, setRef}) => (
  <Box className="menu-button" style={styles.menuButtons}>
    <ReactButton
      conversationIDKey={conversationIDKey}
      ordinal={ordinal}
      showBorder={false}
      tooltipEnabled={false}
    />
    <Box ref={setRef}>
      <Icon
        type="iconfont-ellipsis"
        style={iconCastPlatformStyles(styles.ellipsis)}
        onClick={onClick}
        fontSize={16}
      />
    </Box>
  </Box>
)

const EditedMark = () => (
  <Text type="BodyTiny" style={styles.edited}>
    EDITED
  </Text>
)

const Failure = ({failureDescription, isExplodingUnreadable, onEdit, onRetry, onCancel}) => {
  const error = `${failureDescription}. `
  const resolveByEdit = failureDescription === 'Failed to send: message is too long'
  return isExplodingUnreadable ? (
    <Text type="BodySmall" style={styles.fail}>
      This exploding message is not available to you.
    </Text>
  ) : (
    <Text type="BodySmall">
      <Text type="BodySmall" style={styles.fail}>
        {error}
      </Text>
      {!!onCancel && (
        <Text type="BodySmall" style={styles.failStyleUnderline} onClick={onCancel}>
          Cancel
        </Text>
      )}
      {!!onCancel && <Text type="BodySmall"> or </Text>}
      {!!onEdit &&
        resolveByEdit && (
          <Text type="BodySmall" style={styles.failStyleUnderline} onClick={onEdit}>
            Edit
          </Text>
        )}
      {!!onRetry &&
        !resolveByEdit && (
          <Text type="BodySmall" style={styles.failStyleUnderline} onClick={onRetry}>
            Retry
          </Text>
        )}
    </Text>
  )
}

const sendIndicatorWidth = 32

const LeftSide = props => (
  <Box style={styles.leftSide}>
    {props.includeHeader && <UserAvatar author={props.author} onAuthorClick={props.onAuthorClick} />}
  </Box>
)

const RightSide = props => (
  <Box style={styles.rightSideContainer}>
    <Box
      style={collapseStyles([styles.rightSide, props.includeHeader && styles.hasHeader])}
      className="message-wrapper"
    >
      {props.includeHeader && (
        <Username
          username={props.author}
          isYou={props.isYou}
          isFollowing={props.isFollowing}
          isBroken={props.isBroken}
          onClick={props.onAuthorClick}
        />
      )}
      <Box style={styles.textContainer} className="message">
        {/* TODO remove the `|| props.isExplodingUnreadable` when a fix for inadvertent error messages is in.
          The problem is that `isExplodingUnreadable` is coming as true without `props.exploded` sometimes.  */}
        <ExplodingHeightRetainer
          explodedBy={props.explodedBy}
          exploding={props.exploding}
          measure={props.measure}
          messageKey={props.messageKey}
          style={styles.flexOneColumn}
          retainHeight={props.exploded || props.isExplodingUnreadable}
        >
          <props.innerClass
            message={props.message}
            isEditing={props.isEditing}
            toggleShowingMenu={props.toggleShowingMenu}
          />
          {props.isEdited && <EditedMark />}
        </ExplodingHeightRetainer>
        {!isMobile &&
          !props.exploded && (
            <MenuButtons
              conversationIDKey={props.conversationIDKey}
              ordinal={props.ordinal}
              setRef={props.setAttachmentRef}
              onClick={props.toggleShowingMenu}
            />
          )}
        <MessagePopup
          attachTo={props.attachmentRef}
          message={props.message}
          onHidden={props.toggleShowingMenu}
          position="top center"
          visible={props.showingMenu}
        />
        {props.isRevoked && (
          <Icon
            type="iconfont-exclamation"
            style={iconCastPlatformStyles(styles.exclamation)}
            color={globalColors.blue}
            fontSize={14}
          />
        )}
      </Box>
      {!!props.failureDescription &&
        !props.exploded && (
          <Failure
            failureDescription={props.failureDescription}
            isExplodingUnreadable={props.isExplodingUnreadable}
            onRetry={props.onRetry}
            onEdit={props.onEdit}
            onCancel={props.onCancel}
          />
        )}
      <Box style={styles.sendIndicatorContainer}>
        {props.isYou && (
          <SendIndicator
            sent={props.messageSent || props.exploded}
            failed={props.messageFailed}
            style={{marginBottom: 2}}
            id={props.timestamp}
          />
        )}
      </Box>
    </Box>
    {props.exploding && (
      <ExplodingMeta
        exploded={props.exploded}
        explodesAt={props.explodesAt}
        messageKey={props.messageKey}
        pending={props.messagePending || props.messageFailed}
        onClick={props.exploded ? null : props.toggleShowingMenu}
      />
    )}
  </Box>
)

class WrapperAuthor extends React.PureComponent<WrapperAuthorProps & OverlayParentProps> {
  componentDidUpdate(prevProps: WrapperAuthorProps) {
    if (this.props.measure) {
      if (this.props.includeHeader !== prevProps.includeHeader) {
        this.props.measure()
      }
    }
  }

  render() {
    const props = this.props
    return (
      <Box
        style={collapseStyles([
          styles.leftRightContainer,
          props.showingMenu && styles.selected,
          props.includeHeader && styles.hasHeader,
        ])}
      >
        <LeftSide {...props} />
        <RightSide {...props} />
      </Box>
    )
  }
}

const styles = styleSheetCreate({
  edited: {backgroundColor: globalColors.white, color: globalColors.black_20_on_white},
  ellipsis: {marginRight: globalMargins.xtiny},
  exclamation: {
    paddingBottom: globalMargins.xtiny,
    paddingTop: globalMargins.xtiny,
  },
  fail: {color: globalColors.red},
  failStyleUnderline: {color: globalColors.red, textDecorationLine: 'underline'},
  flexOneColumn: {...globalStyles.flexBoxColumn, flex: 1},
  flexOneRow: {...globalStyles.flexBoxRow, flex: 1},
  hasHeader: {paddingTop: 6},
  leftRightContainer: {...globalStyles.flexBoxRow, width: '100%'},
  leftSide: platformStyles({
    common: {
      flexShrink: 0,
      marginRight: globalMargins.tiny,
      position: 'relative',
      width: 32,
    },
    isElectron: {
      marginLeft: globalMargins.small,
    },
    isMobile: {
      marginLeft: globalMargins.tiny,
    },
  }),
  menuButtons: {
    ...globalStyles.flexBoxRow,
    alignSelf: 'flex-start',
    alignItems: 'center',
    position: 'relative',
    top: 1,
  },
  rightSide: {
    ...globalStyles.flexBoxColumn,
    flex: 1,
    paddingRight: globalMargins.tiny,
    position: 'relative',
  },
  rightSideContainer: {
    ...globalStyles.flexBoxRow,
    flex: 1,
    paddingBottom: 2,
    paddingRight: globalMargins.tiny,
  },
  selected: {backgroundColor: globalColors.black_05},
  sendIndicator: {marginBottom: 2},
  sendIndicatorContainer: platformStyles({
    common: {
      alignItems: 'flex-start',
      bottom: -2,
      height: 21,
      justifyContent: 'center',
      position: 'absolute',
      right: 0,
      width: sendIndicatorWidth,
    },
    isElectron: {pointerEvents: 'none'},
    isMobile: {
      right: -14,
    },
  }),
  textContainer: {
    ...globalStyles.flexBoxRow,
    borderRadius: 4,
    flex: 1,
  },
  userAvatar: {
    flexShrink: 0,
    height: 32,
    width: 32,
  },
  username: {
    alignSelf: 'flex-start',
    backgroundColor: globalColors.fastBlank,
    marginBottom: 0,
  },
  usernameYou: {
    ...globalStyles.italic,
  },
})

export default WrapperAuthor
